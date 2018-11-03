package actions

import (
	"github.com/harsimranman/accounting/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func setProject(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		p := &models.Project{}
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return missingTransaction
		}
		cid := (c.Value("current_user").(*models.User)).CompanyID
		if err := tx.Where("company_id = ?", cid).Find(p, c.Param("project_id")); err != nil {
			return err
		}
		c.Set("current_project", p)
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
		pid := (c.Value("current_project").(*models.Project)).ID
		if err := tx.Where("project_id = ?", pid).Find(b, c.Param("budget_line_id")); err != nil {
			return err
		}
		c.Set("current_budget_line", b)
		return next(c)
	}
}