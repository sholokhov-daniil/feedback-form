package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/context"
	ex "github.com/sholokhov-daniil/feedback-form/internal/exceptions"
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
// @Success      200 {object} []dto.FormResponse
// @Failure      400 {object} []response.Error
// @Failure      500 {object} []response.Error
// @Router       /forms [get]
// @Security     BearerAuth
func (h *FormHandler) GetAll(w http.ResponseWriter, r *http.Request) {	
	ctx := r.Context()
	u, err := context.GetUser(ctx)

	if err != nil {
		h.handleRepoError(w, err);
		return
	}

	forms, err := h.repo.GetByUserID(ctx, u.ID)

	if err != nil {
		h.handleRepoError(w, err);
		return
	}

	data := normalizer.FormListNormalize(forms)
	h.respondJSON(w, http.StatusOK, data)
}

// Returns a specific user form by its ID
// @Summary Returns a specific user form by its ID
// @Tags         forms
// @Accept       json
// @Produce      json
// @Param        id path string true "Form ID"
// @Success      200 {object} dto.FormResponse
// @Failure      404 {object} []response.Error
// @Failure      500 {object} []response.Error
// @Router       /forms/{id} [get]
func (h *FormHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	form, err := h.repo.GetByID(ctx, id)

	if err != nil {
		h.handleRepoError(w, err);
		return
	}


	data := normalizer.FormNormalize(form)
	h.respondJSON(w, http.StatusOK, data)
}

// respondJSON writes a JSON response with the given status code and data.
// It sets the Content-Type header to application/json and handles encoding errors by logging them.
// This method is used internally by other handler methods to send successful responses.
func (h *FormHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("failed to encode JSON response: %v", err)
    }
}

// respondError writes a JSON error response with the given status code and message.
// It selects the appropriate error response format based on the status code:
// - 404: uses CreateNotFoundErrorResponse
// - 500: uses CreateServerErrorResponse
// - others: falls back to a simple {"error": message} object.
// The actual JSON writing is delegated to respondJSON.
func (h *FormHandler) respondError(w http.ResponseWriter, status int, message string) {
    var resp interface{}

    switch status {
		case http.StatusNotFound:
			resp = response.CreateNotFoundErrorResponse(message)
		case http.StatusInternalServerError:
			resp = response.CreateServerErrorResponse(message)
		default:
			resp = map[string]string{"error": message}
    }

    h.respondJSON(w, status, resp)
}

// handleRepoError processes errors returned by the repository.
// It distinguishes between a "form not found" error (returns HTTP 404)
// and any other unexpected error (logs it and returns HTTP 500).
// The actual HTTP response is sent via respondError.
func (h *FormHandler) handleRepoError(w http.ResponseWriter, err error) {
    if errors.Is(err, ex.ErrFormNotFound) {
        h.respondError(w, http.StatusNotFound, err.Error())
        return
    }

    log.Printf("unexpected repository error: %v", err)
    h.respondError(w, http.StatusInternalServerError, "internal server error")
}