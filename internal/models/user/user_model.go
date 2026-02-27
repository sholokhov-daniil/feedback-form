package models

import (
	"time"

	model "github.com/sholokhov-daniil/feedback-form/internal/models/form"
)

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Active     bool      `gorm:"default:true;not null" json:"active"`
	Name       string    `gorm:"size:255;not null" json:"name"`
	DateCreate time.Time `gorm:"autoCreateTime;not null" json:"date_create"`
	DateUpdate time.Time `gorm:"autoUpdateTime;not null" json:"date_update"`

	AuthMethods []UserAuth   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"auth_methods,omitempty"`
	Forms       []model.Form `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;" json:"forms,omitempty"`
}

func (User) TableName() string {
	return "users"
}
