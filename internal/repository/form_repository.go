package repository

import (
    "context"

    "github.com/jmoiron/sqlx"
    "github.com/sholokhov-daniil/feedback-form/internal/models"
)

type FormRepository struct {
    db *sqlx.DB
}

//
// Создает новый репозиторий 
//
func NewFormRepository(db *sqlx.DB) *FormRepository {
    return &FormRepository{db: db}
}

//
// Создает новую форму пользователя
//
func (r *FormRepository) CreateForm(ctx context.Context, form *models.Form) error {
    query := `
        INSERT INTO forms (id, active, user_id, date_create, date_update)
        VALUES ($1, $2, $3, NOW(), NOW())
    `
    _, err := r.db.ExecContext(ctx, query, form.ID, form.Active, form.UserID)
    return err
}

//
// Возвращаем все формы пользователя
//
func (r *FormRepository) GetFormsByUserID(ctx context.Context, userID int) ([]models.Form, error) {
    query := `
        SELECT id, active, user_id, date_create, date_update
        FROM forms
        WHERE user_id = $1
        ORDER BY date_create DESC
    `

    var forms []models.Form
    err := r.db.SelectContext(ctx, &forms, query, userID)
	
    if err != nil {
        return nil, err
    }

    return forms, nil
}