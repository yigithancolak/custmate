package postgresdb

import "github.com/yigithancolak/custmate/graph/model"

func (s *Store) CreatePaymentTx(orgID string, input *model.CreatePaymentInput) (*model.Payment, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, ErrBeginTransaction
	}

	payment, err := s.CreatePayment(tx, orgID, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	updateCustomerInput := &model.UpdateCustomerInput{
		LastPayment: &payment.Date,
		NextPayment: &input.NextPaymentDate,
	}

	_, err = s.UpdateCustomer(tx, input.CustomerID, updateCustomerInput)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return payment, nil
}
