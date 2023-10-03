package postgresdb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreatePayment(orgID string, input *model.CreatePaymentInput) (*model.Payment, error) {
	query := `INSERT INTO payments (id, amount, payment_type, currency, date, customer_id, org_group_id, organization_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, amount, payment_type, currency, date`

	paymentID := uuid.New().String()

	var payment model.Payment

	err := s.DB.QueryRow(query, paymentID, input.Amount, input.PaymentType, input.Currency, input.Date, input.CustomerID, input.GroupID, orgID).Scan(&payment.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (s *Store) UpdatePayment(id string, input *model.UpdatePaymentInput) error {
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

func (s *Store) DeletePayment(id string) error {
	query := "DELETE FROM payments WHERE id = $1"

	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetPaymentByID(id string, includeCustomer bool) (*model.Payment, error) {
	query := "SELECT id, customer_id, amount, payment_type, currency, date FROM payments WHERE id = $1"
	var payment model.Payment
	var customerID string
	err := s.DB.QueryRow(query, id).Scan(&payment.ID, &customerID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
	if err != nil {
		return nil, err
	}

	if includeCustomer {
		payment.Customer, err = s.GetCustomerByID(customerID, false)
	}

	return &payment, err
}

func (s *Store) ListPaymentsByOrganizationID(orgID string, offset *int, limit *int, startDate string, endDate string) ([]*model.Payment, error) {
	query := `SELECT id, amount, payment_type, currency, date 
	FROM payments 
	WHERE organization_id = $1 
	AND date BETWEEN $2 AND $3 
	LIMIT $4 OFFSET $5`

	rows, err := s.DB.Query(query, orgID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	if rows.Next() {
		var payment model.Payment
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}

func (s *Store) ListPaymentsByGroupID(groupID string, offset *int, limit *int, startDate string, endDate string) ([]*model.Payment, error) {
	query := `
	SELECT id, amount, payment_type, currency, date FROM payments
	WHERE org_group_id = $1
	AND date BETWEEN $2 AND $3 
	LIMIT $4 OFFSET $5
	`
	rows, err := s.DB.Query(query, groupID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	if rows.Next() {
		var payment model.Payment
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}

func (s *Store) ListPaymentsByCustomerID(customerID string, offset *int, limit *int, startDate string, endDate string) ([]*model.Payment, error) {
	query := `
	SELECT id, amount, payment_type, currency, date FROM payments
	WHERE customer_id = $1
	AND date BETWEEN $2 AND $3 
	LIMIT $4 OFFSET $5
	`
	rows, err := s.DB.Query(query, customerID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	if rows.Next() {
		var payment model.Payment
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}

func (s *Store) ListPaymentsByFieldID(fieldName string, ID string, offset *int, limit *int, startDate string, endDate string) ([]*model.Payment, error) {
	query := fmt.Sprintf(`
	SELECT id, amount, payment_type, currency, date FROM payments
	WHERE %s = $1
	AND date BETWEEN $2 AND $3 
	LIMIT $4 OFFSET $5
	`, fieldName)
	rows, err := s.DB.Query(query, ID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	for rows.Next() {
		var payment model.Payment
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.PaymentType, &payment.Currency, &payment.Date)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}
