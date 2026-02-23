package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/context"
	"github.com/sholokhov-daniil/feedback-form/internal/handler/normalizer"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)


// Returns all available forms
// @Summary      List feedback forms
// @Description  Returns all available forms
// @Tags         forms
// @Accept       json
// @Produce      json
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /forms [get]
// @Security     BearerAuth
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

	json.NewEncoder(w).Encode(res)
}


// Returns a specific user form by its ID
// @Summary Returns a specific user form by its ID
// @Tags         forms
// @Accept       json
// @Produce      json
// @Param        id path string true "Form ID"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /form/{id} [get]
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