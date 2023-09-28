package postgresdb

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type PaymentStore struct {
	DB *sqlx.DB
}

func NewPaymentStore(db *sqlx.DB) *PaymentStore {
	return &PaymentStore{
		DB: db,
	}
}

func (s *PaymentStore) CreatePayment(input *model.CreatePaymentInput) (*model.Payment, error) {
	query := `INSERT INTO payments (id, customer_id, amount, payment_type, currency, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, customer_id, amount, payment_type, currency, date`

	paymentID := uuid.New().String()
	payment := model.Payment{
		Customer: &model.Customer{},
	}

	err := s.DB.QueryRow(query, paymentID, input.CustomerID, input.Amount, input.PaymentType, input.Currency, input.Date).Scan(&payment.ID, &payment.Customer.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
