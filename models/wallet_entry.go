package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type WalletEntry struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CompanyID   uuid.UUID `json:"company_id" db:"company_id"`
	Notes       string    `json:"notes" db:"notes"`
	Amount      float64   `json:"amount" db:"amount"`
	PaymentDate time.Time `json:"payment_date" db:"payment_date"`
	PaymentType string    `json:"payment_type" db:"payment_type"`
	ReceiptID   uuid.UUID `json:"receipt_id" db:"receipt_id"`
	Receipt     Receipt   `belongs_to:"receipt"`
}

// String is not required by pop and may be deleted
func (w WalletEntry) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// WalletEntries is not required by pop and may be deleted
type WalletEntries []WalletEntry

// String is not required by pop and may be deleted
func (w WalletEntries) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *WalletEntry) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Notes, Name: "Notes"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *WalletEntry) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *WalletEntry) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
