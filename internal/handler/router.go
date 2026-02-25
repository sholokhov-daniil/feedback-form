package handler

import (
	"net/http"
	"github.com/sholokhov-daniil/feedback-form/internal/middleware"
)

type Route struct {
	Method string
	Path string
	Handler http.HandlerFunc
	Middlewares []middleware.Middleware
}

func Registration(mux *http.ServeMux, routes []Route) {
    for _, route := range routes {
        handler := middleware.Chain(
            http.HandlerFunc(route.Handler),
            route.Middlewares...,
        )
        mux.Handle(route.Method+" "+route.Path, handler)
    }
}