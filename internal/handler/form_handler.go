package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/handler/dto"
)

func GetAllForms(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	
	// Берем данные из BD
	forms := []dto.Form{
		{
			ID:   "1",
			Name: "Пример формы",
			Fields: []dto.FormField{
				{Name: "email", Type: ""},
				{Name: "message", Type: "text"},
			},
		},
		{ID: "2", Name: "Регистрация на событие"},
	}

	json.NewEncoder(response).Encode(forms)
}

func GetFormById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	id := request.PathValue("id")

	// Заглушка для примера
	form := dto.Form{
		ID:   id,
		Name: "Пример формы",
		Fields: []dto.FormField{
			{Name: "email", Type: ""},
			{Name: "message", Type: "text"},
		},
	}

	json.NewEncoder(response).Encode(form)
}