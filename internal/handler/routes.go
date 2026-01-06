package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /forms", GetAllForms)
	mux.HandleFunc("GET /forms/{id}", GetFormById)
}