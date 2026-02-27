package repository

import (
	"context"

	"github.com/sholokhov-daniil/feedback-form/internal/models/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{db: ServiceContainer().Database}
}

func (r *userRepositoryImpl) GetByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User

	res := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&u)

	return &u, res.Error
}