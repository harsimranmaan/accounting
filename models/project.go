package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Project struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	Name        string      `json:"name" db:"name"`
	IsActive    bool        `json:"is_active" db:"is_active"`
	CompanyID   uuid.UUID   `json:"company_id" db:"company_id"`
	Company     Company     `belongs_to:"company"`
	BudgetLines BudgetLines `has_many:"budget_lines"`
}

// String is not required by pop and may be deleted
func (p Project) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}
func (p Project) SelectValue() interface{} {
	return p.ID
}
func (p Project) SelectLabel() string {
	return p.Name
}

// Projects is not required by pop and may be deleted
type Projects []Project

// String is not required by pop and may be deleted
func (p Projects) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Project) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// Budget returns the total allocated budget for a project
func (p *Project) Budget() float64 {
	budgetLine := &BudgetLine{}
	DB.Where("project_id = ?", p.ID).Select("sum(amount) as amount").First(budgetLine)
	return budgetLine.Amount
}

// RemainingBudget returns the remaining allocated budget for a project
func (p *Project) RemainingBudget() float64 {
	return p.Budget() - p.UsedBudget()
}

// UsedBudget returns the budfget that has been used for a project
func (p *Project) UsedBudget() float64 {
	receipt := &Receipt{}
	DB.Where("project_id = ?", p.ID).Select("sum(receipts.amount) as amount").LeftJoin("budget_lines", "receipts.budget_line_id = budget_lines.id").First(receipt)
	return receipt.Amount
}
