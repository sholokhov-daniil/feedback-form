package main

import (
	"log"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/db"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
	"github.com/sholokhov-daniil/feedback-form/internal/handler"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
)


func main() {
	container := loadContainer()

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, container.Database)
	
	if err := http.ListenAndServe(":" + container.Config.Host.Port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func loadContainer() *repository.Container {
	container := repository.ServiceContainer()
	container.Config = *config.Load();

	db := db.Open(&container.Config.DB)
	defer db.Close()

	container.Database = db;

	return container;
}