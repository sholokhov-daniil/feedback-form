package handler

import (
	controller "github.com/sholokhov-daniil/feedback-form/internal/controller/form"
	"github.com/sholokhov-daniil/feedback-form/internal/middleware"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
)

func RouteList() []Route {
	formHandler := newFormHandler()

	return []Route{
		{
			Method: "GET",
			Path: "/api/v1/forms",
			Handler: formHandler.GetAll,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
				middleware.AuthBearerMiddleware,
			},
		},
		{
			Method: "GET",
			Path: "/api/v1/form/{id}",
			Handler: formHandler.GetById,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
				middleware.AuthBearerMiddleware,
			},
		},
	}
}

func newFormHandler() *controller.FormHandler {
	formRepo := repository.NewFormRepository();
	return controller.NewFormHandler(formRepo)
}