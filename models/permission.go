package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Permission struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	ObjectType int       `json:"object_type" db:"object_type"`
	ObjectID   uuid.UUID `json:"object_id" db:"object_id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	GroupID    uuid.UUID `json:"group_id" db:"group_id"`
	CompanyID  uuid.UUID `json:"company_id" db:"company_id"`
	Permitted  bool      `json:"permitted" db:"permitted"`
}

// String is not required by pop and may be deleted
func (p Permission) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Permissions is not required by pop and may be deleted
type Permissions []Permission

// String is not required by pop and may be deleted
func (p Permissions) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Permission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.ObjectType, Name: "ObjectType"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Permission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Permission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
