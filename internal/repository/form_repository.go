package repository

import (
    "context"

    "gorm.io/gorm"
    "github.com/sholokhov-daniil/feedback-form/internal/models"
)

type FormRepository struct {
    db *gorm.DB
}

//
// Создает новый репозиторий 
//
func NewFormRepository(db *gorm.DB) *FormRepository {
    return &FormRepository{db: db}
}

//
// Создает новую форму пользователя
//
func (r *FormRepository) CreateForm(ctx context.Context, form *models.Form) error {
    // GORM автоматически заполнит DateCreate и DateUpdate благодаря тегам autoCreateTime/autoUpdateTime
    result := r.db.WithContext(ctx).Create(form)
    return result.Error
}

//
// Возвращаем все формы пользователя
//
func (r *FormRepository) GetFormsByUserID(ctx context.Context, userID int) ([]models.Form, error) {
    var forms []models.Form
    
    // GORM запрос с условиями и сортировкой
    result := r.db.WithContext(ctx).
        Where("user_id = ?", userID).
        Order("date_create DESC").
        Find(&forms)
    
    return forms, result.Error
}

//
// Получить форму по ID с полями (опционально)
//
func (r *FormRepository) GetFormByID(ctx context.Context, formID string) (*models.Form, error) {
    var form models.Form
    
    result := r.db.WithContext(ctx).
        Preload("Fields"). // Предзагрузка связанных полей
        First(&form, "id = ?", formID)
    
    return &form, result.Error
}

//
// Обновить форму
//
func (r *FormRepository) UpdateForm(ctx context.Context, form *models.Form) error {
    // GORM автоматически обновит DateUpdate благодаря тегу autoUpdateTime
    result := r.db.WithContext(ctx).Save(form)
    return result.Error
}

//
// Удалить форму (мягкое удаление если есть DeletedAt, иначе физическое)
//
func (r *FormRepository) DeleteForm(ctx context.Context, formID string) error {
    result := r.db.WithContext(ctx).
        Delete(&models.Form{}, "id = ?", formID)
    return result.Error
}

//
// Получить форму по ID и UserID (для проверки владельца)
//
func (r *FormRepository) GetFormByIDAndUserID(ctx context.Context, formID string, userID int) (*models.Form, error) {
    var form models.Form
    
    result := r.db.WithContext(ctx).
        Where("id = ? AND user_id = ?", formID, userID).
        First(&form)
    
    return &form, result.Error
}
