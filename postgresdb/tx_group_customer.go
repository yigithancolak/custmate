package postgresdb

import (
	"fmt"

	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateCustomerWithTx(input *model.CreateCustomerInput, organizationID string) (*model.Customer, error) {

	tx, err := s.DB.Begin()
	if err != nil {
		return nil, ErrBeginTransaction
	}

	customer, err := s.Customers.CreateCustomer(tx, organizationID, input)
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

func (s *Store) UpdateCustomerWithTx(customerID string, input *model.UpdateCustomerInput) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return ErrBeginTransaction
	}

	_, err = s.Customers.UpdateCustomer(tx, customerID, input)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return ErrRollbackTransaction
		}
		return err
	}

	if input.Groups != nil {
		err = s.Customers.UpdateCustomerGroupsAssociations(tx, customerID, input.Groups)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return ErrRollbackTransaction
			}
			return err
		}
		//TODO: ADD GROUPS TO RETURN OBJECT
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return ErrRollbackTransaction
		}
		return err
	}

	return nil
}
