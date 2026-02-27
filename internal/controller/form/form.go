package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sholokhov-daniil/feedback-form/internal/context"
	ex "github.com/sholokhov-daniil/feedback-form/internal/exceptions"
	"github.com/sholokhov-daniil/feedback-form/internal/handler/dto"
	"github.com/sholokhov-daniil/feedback-form/internal/handler/normalizer"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)

type FormHandler struct {
	repo repository.FormRepository
}

type okCreate struct {
	ID string `json:"id"`
}

func NewFormHandler(repo repository.FormRepository) *FormHandler {
	return &FormHandler{repo: repo}
}

// Creates a new form
// @Summary      Creates a new form
// @Description  Creates a new form for a specific user
// @Tags         forms
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateFormRequest true "Form creation data"
// @Success      200 {object} dto.FormResponse
// @Failure      400 {array} response.Error
// @Failure      500 {array} response.Error
// @Router       /forms [post]
// @Security     BearerAuth
func (h *FormHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	u, err := context.GetUser(ctx)

	if err != nil {
		h.handleRepoError(w, err);
		return
	}

	var req dto.CreateFormRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid JSON format")
        return
	}

	model := req.ToModel(u.ID)

	if err := h.repo.Create(ctx, &model); err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := dto.ToFormResponse(model)
	w.Header().Set("Location", "/forms/" + model.ID)

	h.respondJSON(w, http.StatusCreated, response);
}

// Returns all available forms
// @Summary      List feedback forms
// @Description  Returns all available forms
// @Tags         forms
// @Accept       json
// @Produce      json
// @Success      200 {array} dto.FormResponse
// @Failure      400 {array} response.Error
// @Failure      500 {array} response.Error
// @Router       /forms [get]
// @Security     BearerAuth
func (h *FormHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
// @Failure      404 {array} response.Error
// @Failure      500 {array} response.Error
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
		case http.StatusBadRequest:
			resp = response.CreateBadRequestError(message)
		default:
			resp = response.Error{
				Message: message,
				Code: "",
			}
    }

    h.respondJSON(w, status, resp)
}

// handleRepoError processes errors returned by the repository.
// It distinguishes between a "form not found" error (returns HTTP 404)
// and any other unexpected error (logs it and returns HTTP 500).
// The actual HTTP response is sent via respondError.
func (h *FormHandler) handleRepoError(w http.ResponseWriter, err error) {
    if errors.Is(err, ex.ErrorFormNotFound) {
        h.respondError(w, http.StatusNotFound, err.Error())
        return
    }

    log.Printf("unexpected repository error: %v", err)
    h.respondError(w, http.StatusInternalServerError, "internal server error")
}