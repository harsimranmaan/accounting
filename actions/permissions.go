package actions

import (
	"accounting/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Permission)
// DB Table: Plural (permissions)
// Resource: Plural (Permissions)
// Path: Plural (/permissions)
// View Template Folder: Plural (/templates/permissions/)

// PermissionsResource is the resource for the Permission model
type PermissionsResource struct {
	buffalo.Resource
}

// List gets all Permissions. This function is mapped to the path
// GET /permissions
func (v PermissionsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	permissions := &models.Permissions{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Permissions from the DB
	if err := q.All(permissions); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, permissions))
}

// Show gets the data for one Permission. This function is mapped to
// the path GET /permissions/{permission_id}
func (v PermissionsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Permission
	permission := &models.Permission{}

	// To find the Permission the parameter permission_id is used.
	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, permission))
}

// New renders the form for creating a new Permission.
// This function is mapped to the path GET /permissions/new
func (v PermissionsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Permission{}))
}

// Create adds a Permission to the DB. This function is mapped to the
// path POST /permissions
func (v PermissionsResource) Create(c buffalo.Context) error {
	// Allocate an empty Permission
	permission := &models.Permission{}

	// Bind permission to the html form elements
	if err := c.Bind(permission); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(permission)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, permission))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Permission was created successfully")

	// and redirect to the permissions index page
	return c.Render(201, r.Auto(c, permission))
}

// Edit renders a edit form for a Permission. This function is
// mapped to the path GET /permissions/{permission_id}/edit
func (v PermissionsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Permission
	permission := &models.Permission{}

	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, permission))
}

// Update changes a Permission in the DB. This function is mapped to
// the path PUT /permissions/{permission_id}
func (v PermissionsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Permission
	permission := &models.Permission{}

	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Permission to the html form elements
	if err := c.Bind(permission); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(permission)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, permission))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Permission was updated successfully")

	// and redirect to the permissions index page
	return c.Render(200, r.Auto(c, permission))
}

// Destroy deletes a Permission from the DB. This function is mapped
// to the path DELETE /permissions/{permission_id}
func (v PermissionsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Permission
	permission := &models.Permission{}

	// To find the Permission the parameter permission_id is used.
	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(permission); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Permission was destroyed successfully")

	// Redirect to the permissions index page
	return c.Render(200, r.Auto(c, permission))
}
