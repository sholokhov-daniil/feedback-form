package repository

import (
    "errors"
    "context"

    "gorm.io/gorm"
    ex "github.com/sholokhov-daniil/feedback-form/internal/exceptions"
    "github.com/sholokhov-daniil/feedback-form/internal/models/form"
)

type FormRepository interface {
    Create(ctx context.Context, form *models.Form) error
    GetByID(ctx context.Context, formID string) (*models.Form, error)
    GetByUserID(ctx context.Context, userID int) ([]models.Form, error)
    GetByIDAndUserID(ctx context.Context, formID string, userID int) (*models.Form, error)
}

type formRepositoryImpl struct {
    db *gorm.DB
}

//
// Создает новый репозиторий 
//
func NewFormRepository() FormRepository {
    db := ServiceContainer().Database
    return &formRepositoryImpl{db: db}
}

//
// Создает новую форму пользователя
//
func (r *formRepositoryImpl) Create(ctx context.Context, form *models.Form) error {
    result := r.db.WithContext(ctx).Create(form)
    return result.Error
}

//
// Возвращаем все формы пользователя
//
func (r *formRepositoryImpl) GetByUserID(ctx context.Context, u int) ([]models.Form, error) {
    var forms []models.Form
    
    // GORM запрос с условиями и сортировкой
    result := r.db.WithContext(ctx).
        Where("user_id = ?", u).
        Order("date_create DESC").
        Find(&forms)
    
    return forms, result.Error
}

//
// Получить форму по ID с полями (опционально)
//
func (r *formRepositoryImpl) GetByID(ctx context.Context, formID string) (*models.Form, error) {
    var form models.Form
    
    res := r.db.WithContext(ctx).
        Preload("Fields"). // Предзагрузка связанных полей
        First(&form, "id = ?", formID)
    
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, ex.ErrorFormNotFound
	}

    return &form, res.Error
}

//
// Получить форму по ID и UserID (для проверки владельца)
//
func (r *formRepositoryImpl) GetByIDAndUserID(ctx context.Context, formID string, userID int) (*models.Form, error) {
    var form models.Form
    
    result := r.db.WithContext(ctx).
        Where("id = ? AND user_id = ?", formID, userID).
        First(&form)
    
    return &form, result.Error
}
