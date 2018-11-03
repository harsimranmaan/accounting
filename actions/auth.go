package actions

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/harsimranman/accounting/models"

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

// AuthCallback handles redirects from the third-party auth
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
	c.Flash().Add("success", "You are successfully logged in")
	return c.Redirect(302, "/")
}

func setUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx, ok := c.Value("tx").(*pop.Connection)
			if !ok {
				return missingTransaction
			}
			if err := tx.Eager("Company").Find(u, uid); err != nil {
				return err
			}
			c.Set("current_user", u)
			// tx = tx.Where("company_id = ?", u.CompanyID)
			// c.Set("tx", tx)
		}
		return next(c)
	}
}

func authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if u := c.Value("current_user"); u == nil {
			c.Flash().Add("danger", "You are not authorized to perform this action. Ask you admin for permission")
			return c.Redirect(302, "/")
		}
		//req:= c.Request()
		// user:= c.Value("current_user").(*models.User)
		// roleCheck:= fmt.Sprintf("%s_%s",req.URL.Path, req.Method)
		return next(c)
	}
}

func logout(c buffalo.Context) error {
	c.Session().Clear()
	if err := c.Session().Save(); err != nil {
		return err
	}
	c.Flash().Add("success", "You have been logged out")
	return c.Redirect(302, "/")
}
