package layers

import (
	"github.com/Zargerion/url_shortener/internal/model"
	"github.com/Zargerion/url_shortener/pkg/databases"
	"github.com/Zargerion/url_shortener/pkg/hashtable"
)

type IModel struct {
	UrlModel model.UrlModel
}

func Model(pg databases.PostgresClient, ht *hashtable.HashTableStore, flag *bool) *IModel {
	return &IModel{
		UrlModel: model.NewUrlModel(pg, ht, flag),
	}
}
