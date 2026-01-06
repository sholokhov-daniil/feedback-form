package form

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sholokhov-daniil/feedback-form/model"
)

func All(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	
	// Берем данные из BD
	forms := []model.Form{
		{
			ID:   "1",
			Name: "Пример формы",
			Fields: []model.FormField{
				{Name: "email", Type: ""},
				{Name: "message", Type: "text"},
			},
		},
		{ID: "2", Name: "Регистрация на событие"},
	}

	json.NewEncoder(response).Encode(forms)
}

func Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(request, "id")

	// Заглушка для примера
	form := model.Form{
		ID:   id,
		Name: "Пример формы",
		Fields: []model.FormField{
			{Name: "email", Type: ""},
			{Name: "message", Type: "text"},
		},
	}

	json.NewEncoder(response).Encode(form)
}