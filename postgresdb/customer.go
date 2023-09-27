package postgresdb

import (
	"database/sql"

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
	query := "INSERT INTO customers (id, name, phone_number, organization_id) VALUES ($1, $2, $3, $4)"
	customerId := uuid.New().String()

	var customer model.Customer

	err := s.DB.QueryRow(query, customerId, input.Name, input.PhoneNumber, organizationID).Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
