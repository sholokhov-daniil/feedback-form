package main

import (
    "log"
	"os"
    "net/http"

    "github.com/sholokhov-daniil/feedback-form/route"
)

func main() {
	port := os.Getenv("HOST_PORT")

	if port == "" {
		port = "8080"
	}

	r := route.NewRouter()

	log.Fatal(http.ListenAndServe(":" + port, r))
	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}