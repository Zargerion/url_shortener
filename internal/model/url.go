package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"

	"github.com/Zargerion/url_shortener/pkg/databases"
	"github.com/Zargerion/url_shortener/pkg/hashtable"
	"github.com/Zargerion/url_shortener/pkg/util"
)

type UrlModel interface {
	GetFullUrlByShort(ctx context.Context, url string) (full_url string, err error)
	PostUrlToGetShort(ctx context.Context, url string) (short_url string, err error)
}

type UrlModelImpl struct {
	pg databases.PostgresClient
	ht *hashtable.HashTableStore
	flag *bool
}

func NewUrlModel(pg databases.PostgresClient, ht *hashtable.HashTableStore, flag *bool) UrlModel {
	return &UrlModelImpl{pg: pg, ht: ht, flag: flag}
}

// Получает оригинальный URL из короткого, дергая его из бд или хеш таблицы, в зависимости от флага -d
// при запуске сервера.

func (um *UrlModelImpl) GetFullUrlByShort(ctx context.Context, url string) (full_url string, err error) {
	url = fmt.Sprintf("http://localhost:8080/%s", url)
	if *um.flag {

		full_url, ok := um.ht.GetValueByKey(url)
		if ok {
			return full_url, nil
		}

	} else {

		query := "SELECT original_url FROM url_pairs WHERE short_url = $1"

		err := um.pg.QueryRow(context.Background(), query, url).Scan(&full_url)
		if err != nil {
			return "", err
		}
		return full_url, nil

	}
	return "", errors.New("не удалось получить короткий url")
}

// Создает короткий URL из оригинального, записывая его в бд или хеш таблицу, в зависимости от флага -d
// при запуске сервера.

func (um *UrlModelImpl) PostUrlToGetShort(ctx context.Context, url string) (short_url string, err error) {
	var rund_char string
	original_url := url
	if *um.flag {

		// Тут используется относительно простая кастомная хеш-функция для получения шеша от полного урла.
		// Чтобы избежать колизий используется итеративный перебор хеш-функции и герерация новых хешей,
		// пока не будет получен уникальный.
		url_hash := util.GenerateHash(url)
		for i := 0; i < 100; i++ {
			short_url = fmt.Sprintf("http://localhost:8080/%s", url_hash)
			_, ok := um.ht.GetValueByKey(short_url)
			if ok {
				rund_char = util.RandomSymbol()
				url = url + rund_char
				url_hash = util.GenerateHash(url)
				continue
			}
			break
		}
		um.ht.AddKeyValue(short_url, original_url)
		return short_url, nil

	} else {

		// Тоже самое, только перегенерация короткого урла осуществляется, исходя из ошибки постгреса при добавлении. 
		query := "INSERT INTO url_pairs (short_url, original_url) VALUES ($1, $2)"
		url_hash := util.GenerateHash(url)
		for i := 0; i < 50; i++ {
			short_url = fmt.Sprintf("http://localhost:8080/%s", url_hash)
			_, err := um.pg.Exec(context.Background(), query, short_url, original_url)
			if err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) && pgErr.Code == "23505" { // Код ошибки для нарушения уникальности
					rund_char = util.RandomSymbol()
					url = url + rund_char
					url_hash = util.GenerateHash(url)
					continue
				}
				return "", err
			}
			break
		}
		return short_url, nil
	}
}
