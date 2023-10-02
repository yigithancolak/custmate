package postgresdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateCustomer(tx *sql.Tx, organizationID string, input *model.CreateCustomerInput) (*model.Customer, error) {
	query := "INSERT INTO customers (id, name, phone_number, last_payment, next_payment, organization_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, phone_number, last_payment, next_payment, organization_id"
	customerId := uuid.New().String()

	var customer model.Customer
	var orgID string

	err := tx.QueryRow(query, customerId, input.Name, input.PhoneNumber, input.LastPayment, input.NextPayment, organizationID).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.LastPayment, &customer.NextPayment, &orgID)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("%w: %v", ErrRollbackTransaction, rbErr)
		}
		return nil, err
	}

	return &customer, nil
}

func (s *Store) UpdateCustomer(tx *sql.Tx, customerID string, input *model.UpdateCustomerInput) (*model.Customer, error) {
	baseQuery := "UPDATE customers SET "
	returnQuery := " RETURNING name, phone_number, last_payment, next_payment"
	var updates []string
	var args []interface{}

	idx := 1
	if input.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", idx))
		args = append(args, input.Name)
		idx++
	}

	if input.PhoneNumber != nil {
		updates = append(updates, fmt.Sprintf("phone_number = $%d", idx))
		args = append(args, input.PhoneNumber)
		idx++
	}

	if input.LastPayment != nil {
		updates = append(updates, fmt.Sprintf("last_payment = $%d", idx))
		args = append(args, input.LastPayment)
		idx++
	}

	if input.NextPayment != nil {
		updates = append(updates, fmt.Sprintf("next_payment = $%d", idx))
		args = append(args, input.NextPayment)
		idx++
	}

	if input.Active != nil {
		updates = append(updates, fmt.Sprintf("active = $%d", idx))
		args = append(args, input.Active)
		idx++
	}

	args = append(args, customerID)

	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery

	customer := &model.Customer{
		ID: customerID,
	}

	err := tx.QueryRow(query, args...).Scan(&customer.Name, &customer.PhoneNumber, &customer.LastPayment, &customer.NextPayment)
	if err != nil {
		return nil, err
	}

	return customer, nil

}

func (s *Store) UpdateCustomerGroupsAssociations(tx *sql.Tx, customerID string, groups []*string) error {
	_, err := tx.Exec("DELETE FROM customer_groups WHERE customer_id = $1", customerID)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO customer_groups (customer_id, org_group_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert new associations
	for _, groupID := range groups {
		_, err := stmt.Exec(customerID, groupID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) DeleteCustomer(id string) error {
	query := "DELETE FROM customers WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	return err
}

func (s *Store) GetCustomersByGroupID(id string) ([]*model.Customer, error) {
	query := `SELECT id, name, phone_number, last_payment, next_payment, active FROM customers c
	JOIN customer_groups cg ON c.id = cg.customer_id
	WHERE cg.org_group_id = $1`

	rows, err := s.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.LastPayment, &customer.NextPayment, &customer.Active)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}

func (s *Store) GetCustomerByID(id string, includeGroups bool) (*model.Customer, error) {
	query := "SELECT id, name, phone_number, last_payment, next_payment, active FROM customers WHERE id = $1"

	var customer model.Customer
	err := s.DB.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.LastPayment, &customer.NextPayment, &customer.Active)
	if err != nil {
		return nil, err
	}
	if includeGroups {
		customer.Groups, err = s.ListGroupsByCustomerID(id)
		if err != nil {
			return nil, err
		}
	}

	return &customer, nil
}

func (s *Store) ListCustomersByGroupID(groupID string, offset *int, limit *int) ([]*model.Customer, error) {
	query := `
	SELECT id, name, phone_number, last_payment, next_payment, active FROM customers c
	JOIN customer_groups cg ON c.id = cg.customer_id
	WHERE cg.org_group_id = $1
	LIMIT $2
	OFFSET $3
	`

	rows, err := s.DB.Query(query, groupID, limit, offset)
	if err != nil {
		return nil, err
	}

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.LastPayment, &customer.NextPayment, &customer.Active)
		if err != nil {
			return nil, err
		}

		// NO NEEDED BECAUSE THEY ALL IN SAME GROUP !!!
		// customer.Groups, err = s.ListGroupsByCustomerID(customer.ID)
		// if err != nil {
		// 	return nil, err
		// }

		customers = append(customers, &customer)
	}

	return customers, nil
}
