package databases

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type PostgresClient interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func NewPostgresClient(ctx context.Context) (pool *pgxpool.Pool, err error) {

	maxConn, _ := strconv.ParseInt(os.Getenv("DB_MAX_CONNECTIONS"), 10, 32)
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	maxAttempts, _ := strconv.Atoi(os.Getenv("DB_MAX_ATTEMPTS"))
	maxDelay, _ := strconv.Atoi(os.Getenv("DB_MAX_DELAY"))

	dsn := os.Getenv("DB_SERVER_URL")

	err = DoWithAttempts(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pgxCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse config: %v\n", err)
		}

		pgxCfg.MaxConns = int32(maxConn)
		pgxCfg.MaxConnIdleTime = time.Duration(maxIdleConn) * time.Second
		pgxCfg.MaxConnLifetime = time.Duration(maxLifetimeConn) * time.Second

		pool, err = pgxpool.ConnectConfig(ctx, pgxCfg)
		if err != nil {
			log.Println("Failed to connect to postgres... Going to do the next attempt")

			return err
		}

		return nil
	}, maxAttempts, time.Duration(maxDelay)*time.Second)

	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to postgres")
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}

		return nil
	}

	return err
}
