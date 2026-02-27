package repository

import (
	"time"

	"github.com/sholokhov-daniil/feedback-form/internal/models/user"
)

func GetByToken(token string) (*models.UserAuth, error) {
	db := ServiceContainer().Database

	var ua models.UserAuth

	now := time.Now()

	// Предзагрузка связанного AuthType
	result := db.
		Preload("AuthType").
		Where("secret_hash = ?", token).
		Where("active = ?", true).
		Where(db.Where("expires_at IS NULL").Or("expires_at > ?", now)).
		First(&ua)

	if result.Error != nil {
		return nil, result.Error
	}

	return &ua, nil
}
