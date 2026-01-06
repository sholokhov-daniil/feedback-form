package main

import (
	"log"
	"os"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/handler"
)


func main() {
	port := os.Getenv("HOST_PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)


	log.Fatal(http.ListenAndServe(":" + port, mux))
	if err := http.ListenAndServe(":" + port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}