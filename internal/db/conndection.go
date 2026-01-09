package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
)

func Open(cfg *config.DBConfig) *sqlx.DB {
	if cfg.User == "" || cfg.Password == "" || cfg.Name == "" || cfg.Host == "" {
		log.Fatal("Database environment variables are not set")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	fmt.Println("Database connected")

	return db;
} 
