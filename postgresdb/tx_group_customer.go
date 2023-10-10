package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateCustomerWithTx(input *model.CreateCustomerInput, organizationID string) (*model.Customer, error) {

	tx, err := s.DB.Begin()
	if err != nil {
		return nil, ErrBeginTransaction
	}

	customer, err := s.CreateCustomer(tx, organizationID, input)
	if err != nil {
		return nil, err
	}

	for _, gID := range input.Groups {
		query := "INSERT INTO customer_groups (customer_id,org_group_id) VALUES($1,$2)"
		_, err := tx.Exec(query, customer.ID, gID)
		if err != nil {
			return nil, fmt.Errorf("error inserting into customer_groups: %w", err)
		}
	}

	//TODO: ADD GROUPS ARRAY FETCHING

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	return customer, nil
}

func (s *Store) UpdateCustomerWithTx(customerID string, input *model.UpdateCustomerInput) (*model.Customer, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, ErrBeginTransaction
	}

	customer, err := s.UpdateCustomer(tx, customerID, input)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	if input.Groups != nil {
		groups, err := s.UpdateCustomerGroupsAssociations(tx, customerID, input.Groups)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, ErrRollbackTransaction
			}
			return nil, err
		}

		customer.Groups = groups
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	return customer, nil
}

func (s *Store) UpdateCustomerGroupsAssociations(tx *sql.Tx, customerID string, groupsIDs []*string) ([]*model.Group, error) {
	_, err := tx.Exec("DELETE FROM customer_groups WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO customer_groups (customer_id, org_group_id) VALUES ($1, $2)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Insert new associations
	for _, groupID := range groupsIDs {
		_, err := stmt.Exec(customerID, groupID)
		if err != nil {
			return nil, err
		}
	}

	// Fetch the groups
	var groups []*model.Group
	for _, groupID := range groupsIDs {
		group, err := s.GetGroupByID(*groupID)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}
