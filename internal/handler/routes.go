package handler

import (
	formController "github.com/sholokhov-daniil/feedback-form/internal/controller/form"
	sysController "github.com/sholokhov-daniil/feedback-form/internal/controller"
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
			Path: "/api/v1/forms/{id}",
			Handler: formHandler.GetById,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
				middleware.AuthBearerMiddleware,
			},
		},
		{
			Method: "GET",
			Path: "/",
			Handler: sysController.PageNotFound,
			Middlewares: []middleware.Middleware{
				middleware.JSONMiddleware,
			},
		},
	}
}

func newFormHandler() *formController.FormHandler {
	formRepo := repository.NewFormRepository();
	return formController.NewFormHandler(formRepo)
}