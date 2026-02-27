package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Form represents a form entity owned by a specific user.
//
// A Form is the aggregate root that contains one or more associated fields.
// It defines the structure and configuration of a dynamic form.
//
// Fields:
//   - ID: unique identifier of the form (UUID, primary key)
//   - Active: indicates whether the form is enabled and available for use
//   - UserID: identifier of the user who owns the form
//   - Name: human-readable name of the form
//   - DateCreate: timestamp when the form was created (set automatically)
//   - DateUpdate: timestamp when the form was last updated (set automatically)
//   - Fields: collection of associated Field entities (one-to-many relationship)
//
// Relationships:
//   - One Form has many Fields
//   - Fields are deleted automatically when the Form is deleted (CASCADE)
//
// This struct is mapped to the "forms" database table.
type Form struct {
	ID         string    `gorm:"primaryKey;type:uuid" json:"id"`
	Active     bool      `gorm:"default:true;not null" json:"active"`
	UserID     int       `gorm:"not null" json:"user_id"`
	Name       string    `gorm:"not null" json:"name"`
	DateCreate time.Time `gorm:"autoCreateTime;not null" json:"date_create"`
	DateUpdate time.Time `gorm:"autoUpdateTime;not null" json:"date_update"`

	Fields []Field `gorm:"foreignKey:FormID;constraint:OnDelete:CASCADE;" json:"fields,omitempty"`
}

// BeforeCreate is a GORM hook that runs before inserting a new Form into the database.
//
// It ensures that the form has a valid UUID primary key.
// If the ID is empty, a new UUID is generated automatically.
//
// This guarantees that every Form has a unique identifier before persistence.
func (f *Form) BeforeCreate(db *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the database table name for the Form model.
//
// This overrides GORM's default naming convention.
func (Form) TableName() string {
	return "forms"
}
