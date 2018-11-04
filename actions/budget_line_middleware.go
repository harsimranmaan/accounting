package actions

import (
	"github.com/harsimranman/accounting/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

func setProject(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		p := &models.Project{}
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return missingTransaction
		}
		cid := (c.Value("currentUser").(*models.User)).CompanyID
		if err := tx.Where("company_id = ?", cid).Find(p, c.Param("project_id")); err != nil {
			return err
		}
		c.Set("currentProject", p)
		return next(c)
	}
}

func setNullUUID(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("nullUuid", uuid.UUID{}.String())
		c.Set("TIME_FORMAT", "2006-Jan-02")
		return next(c)
	}
}

func setBudgetLine(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		b := &models.BudgetLine{}
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return missingTransaction
		}
		p := c.Value("currentProject").(*models.Project)
		if err := tx.Where("project_id = ?", p.ID).Find(b, c.Param("budget_line_id")); err != nil {
			return err
		}
		b.Project = *p
		c.Set("currentBudgetLine", b)
		return next(c)
	}
}