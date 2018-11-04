package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Receipt struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	ReceiptName   string     `json:"receipt_name" db:"receipt_name"`
	ReceiptNumber string     `json:"receipt_number" db:"receipt_number"`
	ReceiptType   string     `json:"receipt_type" db:"receipt_type"`
	ReceiptDate   time.Time  `json:"receipt_date" db:"receipt_date"`
	Amount        float64    `json:"amount" db:"amount"`
	BudgetLineID  uuid.UUID  `json:"budget_line_id" db:"budget_line_id"`
	BudgetLine    BudgetLine `belongs_to:"budget_line"`
	CompanyID     uuid.UUID  `json:"company_id" db:"company_id"`
	Company       Company    `belongs_to:"company"`
	//WalletEntry   WalletEntry `has_one:"wallet_entry" fk_id:"receipt_id"`
}

// String is not required by pop and may be deleted
func (r Receipt) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Receipts is not required by pop and may be deleted
type Receipts []Receipt

// String is not required by pop and may be deleted
func (r Receipts) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Receipt) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.ReceiptName, Name: "ReceiptName"},
		&validators.StringIsPresent{Field: r.ReceiptNumber, Name: "ReceiptNumber"},
		&validators.StringIsPresent{Field: r.ReceiptType, Name: "ReceiptType"},
		&validators.TimeIsPresent{Field: r.ReceiptDate, Name: "ReceiptDate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Receipt) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Receipt) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
