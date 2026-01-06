package models

type AuthType struct {
    ID          int     `db:"id" json:"id"`
    Code        string  `db:"code" json:"code"`
    Name        string  `db:"name" json:"name"`
    Description *string `db:"description" json:"description,omitempty"` // nullable
}