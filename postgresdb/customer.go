package postgresdb

import (
	"database/sql"
	"fmt"

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
