package controller

import (
	"net/http"

	"encoding/json"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response.CreateNotFoundErrorResponse("Method not found"))
}
