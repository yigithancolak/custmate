package postgresdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type PaymentTestSuite struct {
	StoreTestSuite
	Organization *model.Organization
	Instructor   *model.Instructor
	Group        *model.Group
	Customer     *model.Customer
	Payment      *model.Payment
}

func TestPaymentSuite(t *testing.T) {
	suite.Run(t, new(PaymentTestSuite))
}

func (s *PaymentTestSuite) SetupTest() {

	s.Organization = s.createRandomOrganization()
	s.Instructor = s.createRandomInstructor(s.Organization.ID)
	s.Group = s.createRandomGroup(s.Instructor.ID, s.Organization.ID)
	s.Customer = s.createRandomCustomer(s.Organization.ID, []string{s.Group.ID})
	s.Payment = s.createRandomPayment(s.Organization.ID, s.Customer.ID, s.Group.ID)
}

func (s *PaymentTestSuite) createRandomPayment(orgID, customerID, groupID string) *model.Payment {
	args := &model.CreatePaymentInput{
		CustomerID: customerID,
		GroupID:    groupID,
		//
		Amount:          util.RandomIntBetween(100, 500),
		Date:            time.Now().Format(dateFormat),
		NextPaymentDate: time.Now().AddDate(0, 1, 0).Format(dateFormat),
		PaymentType:     util.RandomPaymentType(),
		Currency:        util.RandomCurrency(),
	}

	payment, err := s.store.CreatePaymentTx(orgID, args)
	s.NoError(err)
	s.NotNil(payment)

	s.Equal(args.Amount, payment.Amount)
	s.Equal(args.Date, convertDateFormat(payment.Date))
	s.Equal(args.PaymentType, payment.PaymentType)
	s.Equal(args.Currency, payment.Currency)

	customer, err := s.store.GetCustomerByID(customerID, false)
	s.NoError(err)
	s.NotNil(customer)
	s.Equal(args.NextPaymentDate, convertDateFormat(customer.NextPayment))

	return payment

}

func (s *PaymentTestSuite) TestCreatePayment() {
	s.NotNil(s.Payment)
}

func (s *PaymentTestSuite) TestUpdatePayment() {
	amount := util.RandomIntBetween(300, 400)
	currency := util.RandomCurrency()
	date := util.RandomDate()
	paymentType := util.RandomPaymentType()

	args := model.UpdatePaymentInput{
		Amount:      &amount,
		Currency:    &currency,
		Date:        &date,
		PaymentType: &paymentType,
	}

	updatedPayment, err := s.store.UpdatePayment(s.Payment.ID, &args)
	s.NoError(err)
	s.NotNil(updatedPayment)

	s.Equal(amount, updatedPayment.Amount)
	s.Equal(currency, updatedPayment.Currency)
	s.Equal(date, convertDateFormat(updatedPayment.Date))
	s.Equal(paymentType, updatedPayment.PaymentType)
}

func (s *PaymentTestSuite) TestDeletePayment() {
	deleteID := s.Payment.ID
	err := s.store.DeletePayment(deleteID)
	s.NoError(err)

	deletedPayment, err := s.store.GetPaymentByID(deleteID, false)
	s.Error(err)
	s.Nil(deletedPayment)
}

func (s *PaymentTestSuite) createRandomPayments(n int) (payments []*model.Payment, orgID string, groupID string, customerID string) {
	org := s.createRandomOrganization()
	instructor := s.createRandomInstructor(org.ID)
	group := s.createRandomGroup(instructor.ID, org.ID)

	customer := s.createRandomCustomer(org.ID, []string{group.ID})
	for i := 0; i < n; i++ {
		payment := s.createRandomPayment(org.ID, customer.ID, group.ID)
		payments = append(payments, payment)
	}

	orgID = org.ID
	groupID = group.ID
	customerID = customer.ID

	return
}

func (s *PaymentTestSuite) verifyPayments(expected, actual *model.Payment) {
	s.Equal(expected.ID, actual.ID)
	s.Equal(expected.Amount, actual.Amount)
	s.Equal(expected.Currency, actual.Currency)
	s.Equal(expected.Date, actual.Date)
	s.Equal(expected.PaymentType, actual.PaymentType)
}

func (s *PaymentTestSuite) TestListPaymentsByFieldID() {
	//FOR ORGANIZATION
	offset := 0
	limit := 10
	payments, orgID, groupID, customerID := s.createRandomPayments(limit)
	date, err := time.Parse(dateFormat, convertDateFormat(payments[0].Date))
	s.NoError(err)
	s.T().Log()
	startDate := date.AddDate(0, -1, 0).Format(dateFormat)
	endDate := date.AddDate(0, 1, 0).Format(dateFormat)

	foundPayments, _, err := s.store.ListPaymentsByFieldID("organization_id", orgID, &offset, &limit, startDate, endDate)
	s.NoError(err)
	s.NotNil(foundPayments)
	s.NotEmpty(foundPayments)

	paymentsMap := make(map[string]*model.Payment, limit)

	for _, p := range payments {
		paymentsMap[p.ID] = p
	}

	for _, foundPayment := range foundPayments {
		expectedPayment, existing := paymentsMap[foundPayment.ID]
		s.True(existing)
		s.NotNil(expectedPayment)

		s.verifyPayments(expectedPayment, foundPayment)

	}

	//FOR GROUP
	foundPayments, _, err = s.store.ListPaymentsByFieldID("org_group_id", groupID, &offset, &limit, startDate, endDate)
	s.NoError(err)
	s.NotNil(foundPayments)
	s.NotEmpty(foundPayments)

	for _, foundPayment := range foundPayments {
		expectedPayment, existing := paymentsMap[foundPayment.ID]
		s.True(existing)
		s.NotNil(expectedPayment)

		s.verifyPayments(expectedPayment, foundPayment)
	}

	//FOR CUSTOMER
	foundPayments, _, err = s.store.ListPaymentsByFieldID("customer_id", customerID, &offset, &limit, startDate, endDate)
	s.NoError(err)
	s.NotNil(foundPayments)
	s.NotEmpty(foundPayments)
	for _, foundPayment := range foundPayments {
		expectedPayment, existing := paymentsMap[foundPayment.ID]
		s.True(existing)
		s.NotNil(expectedPayment)

		s.verifyPayments(expectedPayment, foundPayment)
	}
}
