package handler

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/context"
	"github.com/sholokhov-daniil/feedback-form/internal/handler/normalizer"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)

func GetAllForms(w http.ResponseWriter, r *http.Request) {	
	ctx := r.Context()
	u, err := context.GetUser(ctx)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		return
	}

	repo := repository.NewFormRepository()

	forms, err := repo.GetByUserID(ctx, u.ID)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		return
	}

	res := response.New(normalizer.FormListNormalize(forms))

	log.Println(res)

	json.NewEncoder(w).Encode(res)
}

func GetFormById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	u, err := context.GetUser(ctx)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		return
	}

	repo := repository.NewFormRepository()

	form, err := repo.GetByIDAndUserID(ctx, id, u.ID)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		return
	}

	res := response.New(normalizer.FormNormalize(form))

	json.NewEncoder(w).Encode(res)
}