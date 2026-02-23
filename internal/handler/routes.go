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

func RegisterRoutes(mux *http.ServeMux) {
	routes := getConfig()

    for _, route := range routes {
        handler := middleware.Chain(
            http.HandlerFunc(route.Handler),
            route.Middlewares...,
        )
        mux.Handle(route.Method+" "+route.Path, handler)
    }
}

func getConfig() []Route {
	return []Route{
		{
			Method: "GET",
			Path: "/api/v1/forms",
			Handler: GetAllForms,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
				middleware.AuthBearerMiddleware,
			},
		},
		{
			Method: "GET",
			Path: "/api/v1/form/{id}",
			Handler: GetFormById,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
				middleware.AuthBearerMiddleware,
			},
		},
	}
}