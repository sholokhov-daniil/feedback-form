package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Field represents a form field entity.
//
// A Field defines a single input element within a Form, such as
// a text input, email input, checkbox, or other supported field types.
//
// Fields:
//   - ID: unique identifier of the field (UUID, primary key)
//   - FormID: identifier of the parent Form (foreign key)
//   - Code: unique code used for internal identification
//   - Active: indicates whether the field is enabled
//   - Name: human-readable name of the field
//   - TypeID: identifier of the field type (foreign key to FieldType)
//   - Settings: optional JSON or text configuration for the field
//   - DateCreate: timestamp when the field was created
//   - DateUpdate: timestamp when the field was last updated
//
// Relationships:
//   - Each Field belongs to one Form
//   - Each Field has one FieldType
//
// This struct is mapped to the "fields" database table.
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
	Form     Form      `gorm:"foreignKey:FormID" json:"-"`
	TypeInfo FieldType `gorm:"foreignKey:TypeID" json:"type_info,omitempty"`
}

// BeforeCreate is a GORM hook that runs before inserting a new Field.
//
// It generates a UUID for the field if the ID is not already set.
func (f *Field) BeforeCreate(db *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the database table name for the Field model.
func (Field) TableName() string {
	return "fields"
}
