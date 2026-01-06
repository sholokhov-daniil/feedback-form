package models

import "time"

type Form struct {
    ID         string    `db:"id" json:"id"`
    Active     bool      `db:"active" json:"active"`
    UserID     int       `db:"user_id" json:"user_id"`
    DateCreate time.Time `db:"date_create" json:"date_create"`
    DateUpdate time.Time `db:"date_update" json:"date_update"`
}

type Field struct {
    ID         string    `db:"id" json:"id"`
    FormID     string    `db:"form_id" json:"form_id"`
    Code       string    `db:"code" json:"code"`
    Active     bool      `db:"active" json:"active"`
    Name       string    `db:"name" json:"name"`
    TypeID     int       `db:"type" json:"type_id"`
    Settings   string    `db:"settings" json:"settings"`
    DateCreate time.Time `db:"date_create" json:"date_create"`
    DateUpdate time.Time `db:"date_update" json:"date_update"`
}

type FieldType struct {
    ID   int    `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
}
