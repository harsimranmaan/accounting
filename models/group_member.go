package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type GroupMember struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	GroupID   uuid.UUID `json:"group_id" db:"group_id"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
}

// String is not required by pop and may be deleted
func (g GroupMember) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// GroupMembers is not required by pop and may be deleted
type GroupMembers []GroupMember

// String is not required by pop and may be deleted
func (g GroupMembers) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *GroupMember) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *GroupMember) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *GroupMember) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
