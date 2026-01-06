package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sholokhov-daniil/feedback-form/controller/form"
	"github.com/sholokhov-daniil/feedback-form/middleware"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recover)

	// Убираем /api, маршруты идут сразу от корня
	router.Route("/forms", func(forms chi.Router) {
		forms.Get("/", form.All)
		forms.Get("/{id}", form.Get)
		//forms.Post("/{id}/submit", form.Submit) // если добавим POST
	})

	return router
}