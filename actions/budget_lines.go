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
// Model: Singular (BudgetLine)
// DB Table: Plural (budget_lines)
// Resource: Plural (BudgetLines)
// Path: Plural (/budget_lines)
// View Template Folder: Plural (/templates/budget_lines/)

// BudgetLinesResource is the resource for the BudgetLine model
type BudgetLinesResource struct {
	buffalo.Resource
}

// List gets all BudgetLines. This function is mapped to the path
// GET /budget_lines
func (v BudgetLinesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	budgetLines := &models.BudgetLines{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all BudgetLines from the DB
	if err := q.All(budgetLines); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, budgetLines))
}

// Show gets the data for one BudgetLine. This function is mapped to
// the path GET /budget_lines/{budget_line_id}
func (v BudgetLinesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty BudgetLine
	budgetLine := &models.BudgetLine{}

	// To find the BudgetLine the parameter budget_line_id is used.
	if err := tx.Find(budgetLine, c.Param("budget_line_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, budgetLine))
}

// New renders the form for creating a new BudgetLine.
// This function is mapped to the path GET /budget_lines/new
func (v BudgetLinesResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.BudgetLine{}))
}

// Create adds a BudgetLine to the DB. This function is mapped to the
// path POST /budget_lines
func (v BudgetLinesResource) Create(c buffalo.Context) error {
	// Allocate an empty BudgetLine
	budgetLine := &models.BudgetLine{}

	// Bind budgetLine to the html form elements
	if err := c.Bind(budgetLine); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(budgetLine)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, budgetLine))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "BudgetLine was created successfully")

	// and redirect to the budget_lines index page
	return c.Render(201, r.Auto(c, budgetLine))
}

// Edit renders a edit form for a BudgetLine. This function is
// mapped to the path GET /budget_lines/{budget_line_id}/edit
func (v BudgetLinesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty BudgetLine
	budgetLine := &models.BudgetLine{}

	if err := tx.Find(budgetLine, c.Param("budget_line_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, budgetLine))
}

// Update changes a BudgetLine in the DB. This function is mapped to
// the path PUT /budget_lines/{budget_line_id}
func (v BudgetLinesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty BudgetLine
	budgetLine := &models.BudgetLine{}

	if err := tx.Find(budgetLine, c.Param("budget_line_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind BudgetLine to the html form elements
	if err := c.Bind(budgetLine); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(budgetLine)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, budgetLine))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "BudgetLine was updated successfully")

	// and redirect to the budget_lines index page
	return c.Render(200, r.Auto(c, budgetLine))
}

// Destroy deletes a BudgetLine from the DB. This function is mapped
// to the path DELETE /budget_lines/{budget_line_id}
func (v BudgetLinesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty BudgetLine
	budgetLine := &models.BudgetLine{}

	// To find the BudgetLine the parameter budget_line_id is used.
	if err := tx.Find(budgetLine, c.Param("budget_line_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(budgetLine); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "BudgetLine was destroyed successfully")

	// Redirect to the budget_lines index page
	return c.Render(200, r.Auto(c, budgetLine))
}