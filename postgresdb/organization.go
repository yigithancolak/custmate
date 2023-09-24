package postgresdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type OrganizationStore struct {
	DB *sqlx.DB
}

func NewOrganizationStore(db *sqlx.DB) *OrganizationStore {
	return &OrganizationStore{
		DB: db,
	}
}

func (s *OrganizationStore) CreateOrganization(org *model.Organization, password string) error {
	query := `INSERT INTO organizations (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(query, org.ID, org.Name, org.Email, password)
	return err
}

func (s *OrganizationStore) UpdateOrganization(orgID string, input model.UpdateOrganizationInput) (*model.Organization, error) {
	baseQuery := "UPDATE organizations SET "
	returnQuery := " RETURNING id, name, email"

	// slices to build the dynamic part of the query
	var updates []string
	var args []interface{}

	idx := 1 // Placeholder index counter

	if input.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", idx))
		args = append(args, *input.Name)
		idx++
	}

	if input.Email != nil {
		updates = append(updates, fmt.Sprintf("email = $%d", idx))
		args = append(args, *input.Email)
		idx++
	}

	// Check if any field was provided to update
	if len(updates) == 0 {
		return nil, errors.New("no fields provided to update")
	}

	args = append(args, orgID)
	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery

	var org model.Organization
	err := s.DB.QueryRow(query, args...).Scan(&org.ID, &org.Name, &org.Email)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (s *OrganizationStore) DeleteOrganization(id string) error {
	query := `DELETE FROM organizations WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}
