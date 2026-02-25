package main

import (
	"log"
	"net/http"


	"github.com/sholokhov-daniil/feedback-form/internal/docs"
	"github.com/sholokhov-daniil/feedback-form/internal/config"
	"github.com/sholokhov-daniil/feedback-form/internal/db"
	"github.com/sholokhov-daniil/feedback-form/internal/handler"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
)

// @title          Feedback form API
// @version        1.0
// @description    API for interacting with feedback forms
// @termsOfService http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://t.me/sholokhov22
// @contact.email  sholokhovdaniil@yandex.ru

// @host           localhost:8080
// @BasePath       /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	container := loadContainer()
	
	db, err := container.Database.DB();

	if err == nil {
		defer db.Close()
	}

	mux := http.NewServeMux()

	docs.RunSwagger(mux)

	handler.Registration(mux, handler.RouteList())
	
	if err := http.ListenAndServe(":" + container.Config.Host.Port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func loadContainer() *repository.Container {
	container := repository.ServiceContainer()
	container.Config = *config.Load();

	db := db.Open(&container.Config.DB)

	container.Database = db;

	return container;
}