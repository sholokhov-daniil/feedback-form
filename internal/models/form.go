package models

import (
	"time"

    
    "github.com/google/uuid"
	"gorm.io/gorm"
)

type Form struct {
    ID         string    `gorm:"primaryKey;type:uuid" json:"id"`
    Active     bool      `gorm:"default:true;not null" json:"active"`
    UserID     int       `gorm:"not null" json:"user_id"`
    DateCreate time.Time `gorm:"autoCreateTime;not null" json:"date_create"`
    DateUpdate time.Time `gorm:"autoUpdateTime;not null" json:"date_update"`
    
    Fields []Field `gorm:"foreignKey:FormID;constraint:OnDelete:CASCADE;" json:"fields,omitempty"`
}

func (f *Form) BeforeCreate(db *gorm.DB) error {
    if f.ID == "" {
        f.ID = uuid.New().String()
    }
    return nil
}

func (Form) TableName() string {
    return "forms"
}

type Field struct {
    ID         string    `gorm:"primaryKey;type:uuid" json:"id"`
    FormID     string    `gorm:"type:uuid;not null" json:"form_id"`
    Code       string    `gorm:"size:100;not null" json:"code"`
    Active     bool      `gorm:"default:true;not null" json:"active"`
    Name       string    `gorm:"size:255;not null" json:"name"`
    TypeID     int       `gorm:"column:type;not null" json:"type_id"`
    Settings   string    `gorm:"type:text" json:"settings"`
    DateCreate time.Time `gorm:"autoCreateTime;not null" json:"date_create"`
    DateUpdate time.Time `gorm:"autoUpdateTime;not null" json:"date_update"`
    
    // Связи (опционально)
    Form     Form       `gorm:"foreignKey:FormID" json:"-"`
    TypeInfo FieldType  `gorm:"foreignKey:TypeID" json:"type_info,omitempty"`
}

func (Field) TableName() string {
    return "fields"
}

type FieldType struct {
    ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name string `gorm:"size:100;unique;not null" json:"name"`
}

func (FieldType) TableName() string {
    return "field_types"
}
