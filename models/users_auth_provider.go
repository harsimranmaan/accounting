package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type UsersAuthProvider struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	AuthProvider string    `json:"auth_provider" db:"auth_provider"`
	ExtID        string    `json:"ext_id" db:"ext_id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (u UsersAuthProvider) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UsersAuthProviders is not required by pop and may be deleted
type UsersAuthProviders []UsersAuthProvider

// String is not required by pop and may be deleted
func (u UsersAuthProviders) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UsersAuthProvider) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.AuthProvider, Name: "AuthProvider"},
		&validators.StringIsPresent{Field: u.ExtID, Name: "ExtID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UsersAuthProvider) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UsersAuthProvider) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
