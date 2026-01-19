package models

import "time"

type User struct {
    ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Active     bool      `gorm:"default:true;not null" json:"active"`
    Name       string    `gorm:"size:255;not null" json:"name"`
    DateCreate time.Time `gorm:"autoCreateTime;not null" json:"date_create"`
    DateUpdate time.Time `gorm:"autoUpdateTime;not null" json:"date_update"`
    
    // Связи (опционально)
    AuthMethods []UserAuth `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"auth_methods,omitempty"`
    Forms       []Form     `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;" json:"forms,omitempty"`
}

type UserAuth struct {
    ID         int        `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID     int        `gorm:"not null" json:"user_id"`
    AuthTypeID int        `gorm:"not null" json:"auth_type_id"`
    Identifier string     `gorm:"size:255;not null" json:"identifier"`
    SecretHash string     `gorm:"type:text;not null" json:"-"`
    Active     bool       `gorm:"default:true;not null" json:"active"`
    ExpiresAt  *time.Time `json:"expires_at,omitempty"` // nullable
    DateCreate time.Time  `gorm:"autoCreateTime;not null" json:"date_create"`
    DateUpdate time.Time  `gorm:"autoUpdateTime;not null" json:"date_update"`
    
    // Связи (опционально)
    User     User     `gorm:"foreignKey:UserID" json:"-"`
    AuthType AuthType `gorm:"foreignKey:AuthTypeID" json:"auth_type,omitempty"`
}

func (UserAuth) TableName() string {
    return "user_auth"
}

func (User) TableName() string {
    return "users"
}