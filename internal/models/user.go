package models

import "time"

type Form struct {
    ID         string    `db:"id" json:"id"`
    Active     bool      `db:"active" json:"active"`
    UserID     int       `db:"user_id" json:"user_id"`
    DateCreate time.Time `db:"date_create" json:"date_create"`
    DateUpdate time.Time `db:"date_update" json:"date_update"`
}

type UserAuth struct {
    ID         int       `db:"id" json:"id"`
    UserID     int       `db:"user_id" json:"user_id"`
    AuthTypeID int       `db:"auth_type_id" json:"auth_type_id"`
    Identifier string    `db:"identifier" json:"identifier"`
    SecretHash string    `db:"secret_hash" json:"-"`
    Active     bool      `db:"active" json:"active"`
    ExpiresAt  *time.Time `db:"expires_at" json:"expires_at,omitempty"` // nullable
    DateCreate time.Time `db:"date_create" json:"date_create"`
    DateUpdate time.Time `db:"date_update" json:"date_update"`
}