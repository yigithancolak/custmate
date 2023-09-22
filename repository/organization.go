package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type OrganizationRepository struct {
	DB *sqlx.DB
}

// GetByID retrieves an organization by its ID
func (r *OrganizationRepository) GetByID(id int) (*model.Organization, error) {
	var org model.Organization
	err := r.DB.Get(&org, "SELECT id, name, email FROM organizations WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// List retrieves all organizations with pagination
func (r *OrganizationRepository) List(offset, limit int) ([]*model.Organization, error) {
	var orgs []*model.Organization
	err := r.DB.Select(&orgs, "SELECT id, name, email FROM organizations LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r *OrganizationRepository) Create(name, email, password string) (*model.Organization, error) {
	// Hash the password using the utility function
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	var org model.Organization
	sql := `
		INSERT INTO organizations (name, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, email`
	err = r.DB.QueryRow(sql, name, email, hashedPassword).Scan(&org.ID, &org.Name, &org.Email)
	if err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *OrganizationRepository) Update(id int, name, email, password string) (*model.Organization, error) {
	var updateFields []string
	var args []interface{}
	var org model.Organization

	args = append(args, id) // Starting with id for WHERE clause
	idx := 2                // SQL argument index starting after id

	if name != "" {
		updateFields = append(updateFields, fmt.Sprintf("name=$%d", idx))
		args = append(args, name)
		idx++
	}

	if email != "" {
		updateFields = append(updateFields, fmt.Sprintf("email=$%d", idx))
		args = append(args, email)
		idx++
	}

	if password != "" {
		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			return nil, err
		}
		updateFields = append(updateFields, fmt.Sprintf("password=$%d", idx))
		args = append(args, hashedPassword)
	}

	if len(updateFields) == 0 {
		return nil, errors.New("no fields provided for update")
	}

	sqlStatement := fmt.Sprintf(
		"UPDATE organizations SET %s WHERE id=$1 RETURNING id, name, email",
		strings.Join(updateFields, ", "),
	)

	err := r.DB.QueryRow(sqlStatement, args...).Scan(&org.ID, &org.Name, &org.Email)
	if err != nil {
		return nil, err
	}

	return &org, nil
}
