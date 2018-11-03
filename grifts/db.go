package grifts

import (
	"github.com/harsimranman/accounting/models"

	"github.com/gobuffalo/pop"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			com := &models.Company{}
			com.Name = "HSM Company"
			com.IsActive = true
			err := tx.Save(com)
			if err != nil {
				return err
			}
			return nil
		})
	})

})
