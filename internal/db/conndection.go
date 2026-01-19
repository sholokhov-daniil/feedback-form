package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/sholokhov-daniil/feedback-form/internal/config"
)

func Open(cfg *config.DBConfig) *gorm.DB {
	if cfg.User == "" || cfg.Password == "" || cfg.Name == "" || cfg.Host == "" {
		log.Fatal("Database environment variables are not set")
	}

	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }


	return db
} 
