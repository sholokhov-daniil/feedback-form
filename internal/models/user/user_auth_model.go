package models

import(
	"time"

	model "github.com/sholokhov-daniil/feedback-form/internal/models"
)

type UserAuth struct {
    ID         int            `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID     int            `gorm:"not null" json:"user_id"`
    AuthTypeID int            `gorm:"not null" json:"auth_type_id"`
    Identifier string         `gorm:"size:255;not null" json:"identifier"`
    SecretHash string         `gorm:"type:text;not null" json:"-"`
    Active     bool           `gorm:"default:true;not null" json:"active"`
    ExpiresAt  *time.Time     `json:"expires_at,omitempty"`
    DateCreate time.Time      `gorm:"autoCreateTime;not null" json:"date_create"`
    DateUpdate time.Time      `gorm:"autoUpdateTime;not null" json:"date_update"`
    
    User       User           `gorm:"foreignKey:UserID" json:"-"`
    AuthType   model.AuthType `gorm:"foreignKey:AuthTypeID" json:"auth_type,omitempty"`
}

func (UserAuth) TableName() string {
    return "user_auth"
}