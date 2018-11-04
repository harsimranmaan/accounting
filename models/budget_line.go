package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type BudgetLine struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Amount    float64   `json:"amount" db:"amount"`
	ProjectID uuid.UUID `json:"project_id" db:"project_id"`
	Project   Project   `belongs_to:"project"`
	Receipts  Receipts  `has_many:"receipts"`
}

// String is not required by pop and may be deleted
func (b BudgetLine) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// BudgetLines is not required by pop and may be deleted
type BudgetLines []BudgetLine

// String is not required by pop and may be deleted
func (b BudgetLines) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *BudgetLine) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: b.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *BudgetLine) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *BudgetLine) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (b *BudgetLine) UsedAmount() float64 {
	r := &Receipt{}
	DB.Where("budget_line_id = ?", b.ID).Select("sum(amount) as amount").First(r)
	return r.Amount
}
func (b *BudgetLine) RemainingAmount() float64 {
	return b.Amount - b.UsedAmount()
}
