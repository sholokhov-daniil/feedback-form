package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sholokhov-daniil/feedback-form/internal/middleware"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, mws ...Middleware) http.Handler {
    for i := len(mws) - 1; i >= 0; i-- {
        h = mws[i](h)
    }
    return h
}

func RegisterRoutes(mux *http.ServeMux, db *sqlx.DB) {
	auth := middleware.Auth(db)

	mux.Handle("GET /forms", auth(http.HandlerFunc(GetAllForms)))

	mux.Handle("GET /forms/{id}", Chain(
		http.HandlerFunc(GetFormById), 
		auth,
	))
}