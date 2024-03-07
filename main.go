package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Zargerion/url_shortener/internal/layers"
	"github.com/Zargerion/url_shortener/internal/routes"
	"github.com/Zargerion/url_shortener/pkg/databases"
	"github.com/Zargerion/url_shortener/pkg/hashtable"
	"github.com/Zargerion/url_shortener/pkg/server"
)

// Подгрузка энва.

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Инициализации контекста, флагов, хештаблицы, пула коннекшенов постгреса, уровней и гина.
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	canUseHashTable := flag.Bool("d", false, "Enable HashTable DB")
	flag.Parse()

	htStore := hashtable.NewHashTableStore()

	pg_pool, err := databases.NewPostgresClient(ctx)
	if err != nil {
		log.Panicf("Error connecting to pg: %v\n", err)
	}

	m := layers.Model(pg_pool, htStore, canUseHashTable)
	c := layers.Controller(m)

	gin := server.NewGin()

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Тут инициализируются роуты. Хотя я бы разбил на группы.
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	gin = routes.Url(gin, c.UrlController)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Запуск сервака.
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	serverURL := os.Getenv("SERVER_URL")

	if err := gin.Run(serverURL); err != nil {
		log.Println(err)
	}
}
