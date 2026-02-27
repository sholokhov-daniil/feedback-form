package models

// FieldType represents the type of a form field.
//
// It defines the behavior and data format of a Field,
// such as text, number, email, checkbox, etc.
//
// Fields:
//   - ID: unique identifier of the field type (primary key)
//   - Name: unique name of the field type
//
// This struct is mapped to the "field_types" table.
type FieldType struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:100;unique;not null" json:"name"`
}

// TableName specifies the database table name for the FieldType model.
func (FieldType) TableName() string {
	return "field_types"
}
