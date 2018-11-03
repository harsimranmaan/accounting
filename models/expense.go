package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Expense struct {
	ID             uuid.UUID     `json:"id" db:"id"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	Name           string        `json:"name" db:"name"`
	DocumentNumber string        `json:"document_number" db:"document_number"`
	DateOfPayment  time.Time     `json:"date_of_payment" db:"date_of_payment"`
	CompanyID      uuid.UUID     `json:"company_id" db:"company_id"`
	Company        Company       `belongs_to:"company"`
	WalletEntries  WalletEntries `has_many:"wallet_entries"`
}

// String is not required by pop and may be deleted
func (e Expense) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Expenses is not required by pop and may be deleted
type Expenses []Expense

// String is not required by pop and may be deleted
func (e Expenses) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Expense) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.DocumentNumber, Name: "DocumentNumber"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Expense) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Expense) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
