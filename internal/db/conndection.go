package db

import (
	"fmt"
	"os"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

// GetDB возвращает singleton подключения к базе
func GetDB() *sqlx.DB {
	once.Do(func() {
		// Загружаем переменные окружения из .env
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using environment variables")
		}

		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_DB")
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")

		if user == "" || password == "" || dbName == "" || host == "" || port == "" {
			log.Fatal("Database environment variables are not set")
		}

		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			user, password, host, port, dbName,
		)

		db, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to DB: %v", err)
		}

		// Настройки пула соединений
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(0)

		instance = db
		fmt.Println("Database connected")
	})

	return instance
}