package normalizer

import (
	"github.com/sholokhov-daniil/feedback-form/internal/handler/dto"
	"github.com/sholokhov-daniil/feedback-form/internal/models/form"
)

func FormListNormalize(l []models.Form) []dto.FormResponse {
	var r []dto.FormResponse

	for _, m := range l {
		r = append(r, FormNormalize(&m))
	}

	return r
}

func FormNormalize(m *models.Form) dto.FormResponse {
	return dto.FormResponse{
		ID:         m.ID,
		Active:     m.Active,
		DateCreate: m.DateCreate,
		DateUpdate: m.DateUpdate,
		Fields:     FieldListNormalize(m.Fields),
	}
}

func FieldListNormalize(l []models.Field) []dto.FieldResponse {
	var r []dto.FieldResponse

	for _, m := range l {
		r = append(r, FieldNormalize(&m))
	}

	return r
}

func FieldNormalize(m *models.Field) dto.FieldResponse {
	return dto.FieldResponse{
		ID:         m.ID,
		FormID:     m.FormID,
		Code:       m.Code,
		Active:     m.Active,
		Name:       m.Name,
		TypeID:     m.TypeID,
		DateCreate: m.DateCreate,
		DateUpdate: m.DateUpdate,
	}
}
