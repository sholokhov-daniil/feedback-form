package form

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	
	// Берем данные из BD
	forms := []map[string]interface{}{
		{"id": "1", "name": "Обратная связь"},
		{"id": "2", "name": "Регистрация на событие"},
	}

	json.NewEncoder(response).Encode(forms)
}

func Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(request, "id")

	// Заглушка для примера
	form := map[string]interface{}{
		"id":   id,
		"name": "Пример формы",
		"fields": []map[string]string{
			{"name": "email", "type": "string"},
			{"name": "message", "type": "text"},
		},
	}

	json.NewEncoder(response).Encode(form)
}