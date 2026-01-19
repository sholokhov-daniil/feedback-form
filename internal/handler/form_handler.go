package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/handler/dto"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
)

func GetAllForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	ctx := r.Context()

	db := repository.ServiceContainer().Database
	repo := repository.CreateFormRepository(db)

	forms, err := repo.GetByUserID(ctx, 1)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateErrorResponse(err.Error(), "500"))
		return
	}

	res := response.New(forms)
	
	json.NewEncoder(w).Encode(res)
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