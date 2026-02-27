package dto

import (
	"time"

	models "github.com/sholokhov-daniil/feedback-form/internal/models/form"
)

type FormResponse struct {
	ID string `json:"id"`
	Active bool `json:"active"`
	Name string `json:"name"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
	Fields []FieldResponse `json:"fields"`
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

type CreateFormRequest struct {
	Active bool                 `json:"active"`
	Name   string               `json:"name"`
	Fields []CreateFieldRequest `json:"fields"`
}

type CreateFieldRequest struct {
	FormID string `json:"form_id"`
	Code string `json:"code"`
	Active bool `json:"active"`
	Name string `json:"name"`
	TypeID int `json:"type_id"`
}

//
// ToModel converts CreateFormRequest (DTO) into a models.Form entity.
//
// This method is used when creating a new form. It:
//
//   - copies the base form properties (Active)
//   - assigns the provided userID as the owner of the form
//   - converts field DTOs (CreateFieldRequest) into models.Field entities
//   - does NOT set IDs for the form or fields — IDs are generated automatically
//     by GORM hooks (BeforeCreate) or at the database level
//
// The returned models.Form is ready to be persisted using repository.Create().
//
// Example:
//
//	model := req.ToModel(userID)
//	err := repo.Create(ctx, &model)
//
func (r CreateFormRequest) ToModel(userID int) models.Form {
	fields := make([]models.Field, 0, len(r.Fields))

	for _, f := range r.Fields {
		fields = append(fields, models.Field{
			Code:   f.Code,
			Active: f.Active,
			Name:   f.Name,
			TypeID: f.TypeID,
		})
	}

	return models.Form{
		Active: r.Active,
		UserID: userID,
		Fields: fields,
	}
}

//
// ToFormResponse converts a models.Form entity into a FormResponse DTO.
//
// This method is used to prepare API responses returned to clients. It:
//
//   - copies all public form properties
//   - converts associated models.Field entities into FieldResponse DTOs
//   - includes creation and update timestamps
//   - returns a structure ready for JSON serialization
//
// This method is typically used after creating, retrieving, or updating a form.
//
// Example:
//
//	form, _ := repo.GetByID(ctx, id)
//	response := dto.ToFormResponse(*form)
//	json.NewEncoder(w).Encode(response)
//
func ToFormResponse(m models.Form) FormResponse {
	fields := make([]FieldResponse, len(m.Fields))

	for i, f := range m.Fields {
		fields[i] = FieldResponse{
			ID:         f.ID,
			FormID:     f.FormID,
			Code:       f.Code,
			Active:     f.Active,
			Name:       f.Name,
			TypeID:     f.TypeID,
			DateCreate: f.DateCreate,
			DateUpdate: f.DateUpdate,
		}
	}

	return FormResponse{
		ID:         m.ID,
		Active:     m.Active,
		Name:       m.Name,
		DateCreate: m.DateCreate,
		DateUpdate: m.DateUpdate,
		Fields:     fields,
	}
}