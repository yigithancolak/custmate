package postgresdb

import (
	"fmt"
	"strings"

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
	payment := &model.Payment{
		Customer: &model.Customer{},
	}

	err := s.DB.QueryRow(query, paymentID, input.CustomerID, input.Amount, input.PaymentType, input.Currency, input.Date).Scan(&payment.ID, &payment.Customer.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentStore) UpdatePayment(id string, input *model.UpdatePaymentInput) error {
	baseQuery := "UPDATE payments SET "
	var updates []string
	var args []interface{}

	idx := 1
	if input.Amount != nil {
		updates = append(updates, fmt.Sprintf("amount = $%d", idx))
		args = append(args, input.Amount)
		idx++
	}

	if input.Currency != nil {
		updates = append(updates, fmt.Sprintf("currency = $%d", idx))
		args = append(args, input.Currency)
		idx++
	}

	if input.PaymentType != nil {
		updates = append(updates, fmt.Sprintf("payment_type = $%d", idx))
		args = append(args, input.PaymentType)
		idx++
	}

	if input.Date != nil {
		updates = append(updates, fmt.Sprintf("date = $%d", idx))
		args = append(args, input.Date)
		idx++
	}

	args = append(args, id)

	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx)

	_, err := s.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *PaymentStore) DeletePayment(id string) error {
	query := "DELETE FROM payments WHERE id = $1"

	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
