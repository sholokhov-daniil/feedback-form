package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	ex "github.com/sholokhov-daniil/feedback-form/internal/exceptions"
	"github.com/sholokhov-daniil/feedback-form/internal/context"
	"github.com/sholokhov-daniil/feedback-form/internal/handler/normalizer"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)

type FormHandler struct {
	repo repository.FormRepository
}

func NewFormHandler(repo repository.FormRepository) *FormHandler {
	return &FormHandler{repo: repo}
}

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
func (h *FormHandler) GetAll(w http.ResponseWriter, r *http.Request) {	
	ctx := r.Context()
	u, err := context.GetUser(ctx)

	if err != nil {
		json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		return
	}

	forms, err := h.repo.GetByUserID(ctx, u.ID)

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
// @Failure      404 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /form/{id} [get]
func (h *FormHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	form, err := h.repo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, ex.ErrFormNotFound) {
			w.WriteHeader(http.StatusNotFound);
			json.NewEncoder(w).Encode(response.CreateNotFoundErrorResponse(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError);
			json.NewEncoder(w).Encode(response.CreateServerErrorResponse(err.Error()))
		}

		return
	}

	res := response.New(normalizer.FormNormalize(form))

	json.NewEncoder(w).Encode(res)
}