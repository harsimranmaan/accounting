package actions

import (
	"accounting/models"
	"database/sql"
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/pkg/errors"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return missingTransaction
	}
	up := &models.UsersAuthProvider{}
	q := tx.Where("auth_provider = ? and ext_id = ?", gu.Provider, gu.UserID)
	var exists = true
	err = q.First(up)
	if err != nil {
		if sql.ErrNoRows != errors.Cause(err) {
			return err
		}
		exists = false
	}

	u := &models.User{}
	if exists {
		err = tx.Where("id = ?", up.UserID).First(u)
		if err != nil {
			return err
		}
	} else {
		u.Name = gu.Name
		u.Email = gu.Email
		u.IsActive = true
		com := &models.Company{}
		if err = tx.First(com); err != nil {
			return err
		}
		u.CompanyID = com.ID
		err = tx.Save(u)
		if err != nil {
			return err
		}
		up.AuthProvider = gu.Provider
		up.ExtID = gu.UserID
		up.UserID = u.ID
		err = tx.Save(up)
		if err != nil {
			return err
		}
	}
	c.Session().Set("current_user_id", u.ID)
	if err := c.Session().Save(); err != nil {
		return err
	}
	c.Flash().Add("success", u.Name)
	return c.Redirect(302, "/")
}

func logout(c buffalo.Context) error {
	c.Session().Clear()
	if err := c.Session().Save(); err != nil {
		return err
	}
	return c.Redirect(302, "/")
}
