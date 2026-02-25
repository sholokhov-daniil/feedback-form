package dto

import "time"

type FormResponse struct {
	ID string `json:"id"`
	Active bool `json:"active"`
	Name string `json:"name"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
	Fields []FieldResponse `json:"fields"`
}

type CreateFormRequest struct {
	Active bool `json:"active"`
	Name string `json:"name"`
}

type FieldResponse struct {
	ID string `json:"id"`
	FormID string `json:"form_id"`
	Code string `json:"code"`
	Active bool `json:"active"`
	Name string `json:"name"`
	TypeID int `json:"type_id"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
}