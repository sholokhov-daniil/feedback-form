package main

import (
	"log"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/db"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
	"github.com/sholokhov-daniil/feedback-form/internal/handler"
)


func main() {
	config := config.Load();

	database := db.Open(&config.DB)
	defer database.Close()

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, database)

	log.Fatal(http.ListenAndServe(":" + config.Host.Port, mux))
	if err := http.ListenAndServe(":" + config.Host.Port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}