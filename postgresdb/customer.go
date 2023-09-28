package postgresdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type CustomerStore struct {
	DB *sqlx.DB
}

func NewCustomerStore(db *sqlx.DB) *CustomerStore {
	return &CustomerStore{
		DB: db,
	}
}

func (s *CustomerStore) CreateCustomer(tx *sql.Tx, organizationID string, input *model.CreateCustomerInput) (*model.Customer, error) {
	query := "INSERT INTO customers (id, name, phone_number, next_payment, organization_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, phone_number, next_payment, organization_id"
	customerId := uuid.New().String()

	var customer model.Customer
	var orgID string

	err := tx.QueryRow(query, customerId, input.Name, input.PhoneNumber, input.NextPayment, organizationID).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.NextPayment, &orgID)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("%w: %v", ErrRollbackTransaction, rbErr)
		}
		return nil, err
	}

	return &customer, nil
}

func (s *CustomerStore) UpdateCustomer(tx *sql.Tx, customerID string, input *model.UpdateCustomerInput) (*model.Customer, error) {
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

func (s *CustomerStore) UpdateCustomerGroupsAssociations(tx *sql.Tx, customerID string, groups []*string) error {
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
