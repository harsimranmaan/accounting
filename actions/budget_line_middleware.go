package actions

import (
	"accounting/models"

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
