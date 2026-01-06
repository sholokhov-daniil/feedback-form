package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sholokhov-daniil/feedback-form/controller/form"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Убираем /api, маршруты идут сразу от корня
	router.Route("/forms", func(forms chi.Router) {
		forms.Get("/", form.GetAll)
		forms.Get("/{id}", form.Get)
		//forms.Post("/{id}/submit", form.Submit) // если добавим POST
	})

	return router
}