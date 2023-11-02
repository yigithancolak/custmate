package postgresdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type CustomerTestSuite struct {
	StoreTestSuite
	Organization *model.Organization
	Customer     *model.Customer
}

func TestCustomerSuite(t *testing.T) {
	suite.Run(t, new(CustomerTestSuite))
}

func (s *CustomerTestSuite) VerifySameCustomers(expected, actual *model.Customer) {
	s.Equal(expected.ID, actual.ID)
	s.Equal(expected.Name, actual.Name)
	s.Equal(expected.PhoneNumber, actual.PhoneNumber)
	s.Equal(expected.LastPayment, actual.LastPayment)
	s.Equal(expected.NextPayment, actual.NextPayment)
}

func (s *CustomerTestSuite) SetupTest() {
	s.Organization = s.createRandomOrganization()
	instructor1 := s.createRandomInstructor(s.Organization.ID)
	instructor2 := s.createRandomInstructor(s.Organization.ID)

	group1 := s.createRandomGroup(instructor1.ID, s.Organization.ID)
	group2 := s.createRandomGroup(instructor2.ID, s.Organization.ID)

	s.Customer = s.createRandomCustomer(s.Organization.ID, []string{group1.ID, group2.ID})

}

func (s *StoreTestSuite) createRandomCustomer(orgId string, groupIDs []string) *model.Customer {

	args := &model.CreateCustomerInput{
		Name:        util.RandomName(),
		PhoneNumber: util.RandomPhoneNumber(9),
		LastPayment: util.RandomDate(),
		NextPayment: util.RandomDate(),
		Groups:      groupIDs,
	}
	customer, err := s.store.CreateCustomerWithTx(args, orgId)
	s.NoError(err)
	s.NotNil(customer)
	s.NotEmpty(customer)

	s.Equal(args.Name, customer.Name)
	s.Equal(args.PhoneNumber, customer.PhoneNumber)
	s.Equal(args.LastPayment, convertDateFormat(customer.LastPayment))
	s.Equal(args.NextPayment, convertDateFormat(customer.NextPayment))

	return customer
}

func (s *CustomerTestSuite) createRandomCustomers(n int, orgId string, groupIDs []string) []*model.Customer {

	var customers []*model.Customer
	for i := 0; i < n; i++ {
		createdCustomer := s.createRandomCustomer(orgId, groupIDs)
		customers = append(customers, createdCustomer)
	}

	return customers
}

func (s *CustomerTestSuite) TestCreateCustomer() {
	s.NotNil(s.Customer)

}

func (s *CustomerTestSuite) TestUpdateCustomer() {
	instructor := s.createRandomInstructor(s.Organization.ID)
	group1 := s.createRandomGroup(instructor.ID, s.Organization.ID)
	group2 := s.createRandomGroup(instructor.ID, s.Organization.ID)

	name := util.RandomName()
	phoneNumber := util.RandomPhoneNumber(9)
	lastPayment := util.RandomDate()
	nextPayment := util.RandomDate()
	active := false

	args := &model.UpdateCustomerInput{
		Name:        &name,
		PhoneNumber: &phoneNumber,
		LastPayment: &lastPayment,
		NextPayment: &nextPayment,
		Active:      &active,
		Groups:      []*string{&group1.ID, &group2.ID},
	}
	customer, err := s.store.UpdateCustomerWithTx(s.Customer.ID, args)
	s.NoError(err)
	s.NotNil(customer)
	s.NotEmpty(customer)

	s.Equal(name, customer.Name)
	s.Equal(phoneNumber, customer.PhoneNumber)
	s.Equal(lastPayment, convertDateFormat(customer.LastPayment))
	s.Equal(nextPayment, convertDateFormat(customer.NextPayment))
	s.Equal(active, *customer.Active) // Assuming customer.Active is also a *bool

	var groupIDs []*string
	for _, group := range customer.Groups {
		groupIDs = append(groupIDs, &group.ID)
	}
	s.ElementsMatch(args.Groups, groupIDs)

}

func (s *CustomerTestSuite) TestGetCustomerByID() {
	customer := s.Customer

	foundCustomer, err := s.store.GetCustomerByID(customer.ID, false)
	s.NoError(err)
	s.NotNil(foundCustomer)
	s.NotEmpty(foundCustomer)

	s.VerifySameCustomers(customer, foundCustomer)

}

func (s *CustomerTestSuite) TestDeleteCustomer() {
	customer := s.Customer

	err := s.store.DeleteCustomer(customer.ID)
	s.NoError(err)

	deletedCustomer, err := s.store.GetCustomerByID(customer.ID, false)
	s.Error(err)
	s.Nil(deletedCustomer)
}

func (s *CustomerTestSuite) TestListCustomersByOrganizationID() {
	instructor := s.createRandomInstructor(s.Organization.ID)
	group := s.createRandomGroup(instructor.ID, s.Organization.ID)
	customers := s.createRandomCustomers(9, s.Organization.ID, []string{group.ID})
	customers = append(customers, s.Customer)
	offset := 0
	limit := 10
	/////

	foundCustomers, err := s.store.ListCustomersByOrganizationID(s.Organization.ID, &offset, &limit)
	s.NoError(err)
	s.NotEmpty(foundCustomers)
	s.NotNil(foundCustomers)
	s.Equal(limit, len(foundCustomers))

	customersMap := make(map[string]*model.Customer)

	for _, c := range customers {
		customersMap[c.ID] = c
	}

	for _, foundCustomer := range foundCustomers {
		expectedCustomer, existing := customersMap[foundCustomer.ID]
		s.True(existing)

		s.VerifySameCustomers(expectedCustomer, foundCustomer)
	}
}

func (s *CustomerTestSuite) TestListCustomersByGroupID() {
	instructor := s.createRandomInstructor(s.Organization.ID)
	group := s.createRandomGroup(instructor.ID, s.Organization.ID)
	offset := 0
	limit := 10
	customers := s.createRandomCustomers(limit, s.Organization.ID, []string{group.ID})

	///

	foundCustomers, err := s.store.ListCustomersByGroupID(group.ID, &offset, &limit, false)
	s.NoError(err)
	s.NotEmpty(foundCustomers)
	s.NotNil(foundCustomers)
	s.Equal(limit, len(foundCustomers))

	customersMap := make(map[string]*model.Customer)

	for _, c := range customers {
		customersMap[c.ID] = c
	}

	for _, foundCustomer := range foundCustomers {
		expectedCustomer, existing := customersMap[foundCustomer.ID]
		s.True(existing)

		s.VerifySameCustomers(expectedCustomer, foundCustomer)
	}

}

func (s *CustomerTestSuite) createMockCustomersForFilterSearch(n int, orgID string, groupIDs []string) []*model.Customer {

	var customers []*model.Customer
	for i := 0; i < n; i++ {

		args := &model.CreateCustomerInput{
			Name:        util.RandomName(),
			PhoneNumber: util.RandomPhoneNumber(8),
			Groups:      groupIDs,
			LastPayment: time.Now().AddDate(0, -1, 0).Format(dateFormat),
			NextPayment: time.Now().AddDate(0, 0, 6).Format(dateFormat), // 6 days later customer must pay
		}

		if i%2 != 0 {
			args.NextPayment = time.Now().AddDate(0, 0, -5).Format(dateFormat) // 5 days passed from next payment

		}

		customer, err := s.store.CreateCustomerWithTx(args, orgID)
		s.NoError(err)
		customers = append(customers, customer)
	}

	return customers
}

func (s *CustomerTestSuite) TestListCustomersWithSearchFilter() {
	//NAME, PHONENUMBER, ACTIVE SEARCH
	searchedName := s.Customer.Name[:len(s.Customer.Name)-1]

	args := model.SearchCustomerFilter{
		Name:        &searchedName,
		PhoneNumber: &s.Customer.PhoneNumber,
		Active:      s.Customer.Active,
	}

	//TODO: INCLUDE GROUPS CHANGE THE FALSE VALUE
	foundCustomers, count, err := s.store.ListCustomersWithSearchFilter(args, s.Organization.ID, nil, nil, false)
	s.NoError(err)
	s.Equal(1, count)
	s.Equal(foundCustomers[0].Name, s.Customer.Name)

	// TEST LATE PAYMENT
	latePayment := true
	org := s.createRandomOrganization()
	instructor := s.createRandomInstructor(org.ID)
	group := s.createRandomGroup(instructor.ID, org.ID)
	offset := 0
	limit := 10

	customers := s.createMockCustomersForFilterSearch(limit, org.ID, []string{group.ID})

	args = model.SearchCustomerFilter{
		LatePayment: &latePayment,
	}

	foundCustomers, count, err = s.store.ListCustomersWithSearchFilter(args, org.ID, &offset, &limit, false)
	s.NoError(err)
	s.Equal(limit/2-limit%2, count) //only i%2 != 0  are late payment
	s.NotNil(foundCustomers)
	s.NotEmpty(foundCustomers)

	customersMap := make(map[string]*model.Customer)

	for _, c := range customers {
		s.T().Log(c.NextPayment)
		customersMap[c.ID] = c
	}

	for _, foundCustomer := range foundCustomers {
		expectedCustomer, existing := customersMap[foundCustomer.ID]
		s.True(existing)
		s.NotNil(expectedCustomer)

		parsedNextPayment, err := time.Parse(dateFormat, convertDateFormat(expectedCustomer.LastPayment))
		s.NoError(err)

		s.True(time.Now().After(parsedNextPayment))
	}

	// TEST UPCOMING PAYMENT
	upcomingPayment := true
	upcomingArgs := model.SearchCustomerFilter{
		UpcomingPayment: &upcomingPayment,
	}

	foundCustomers, count, err = s.store.ListCustomersWithSearchFilter(upcomingArgs, org.ID, &offset, &limit, false)
	s.NoError(err)
	s.Equal(limit/2+limit%2, count) //only i%2 == 0  are late payment
	s.NotNil(foundCustomers)
	s.NotEmpty(foundCustomers)

	for _, foundCustomer := range foundCustomers {
		expectedCustomer, existing := customersMap[foundCustomer.ID]
		s.True(existing)
		s.NotNil(expectedCustomer)

		parsedNextPayment, err := time.Parse(dateFormat, convertDateFormat(expectedCustomer.NextPayment))
		s.NoError(err)

		now := time.Now()
		sevenDaysEarlier := parsedNextPayment.AddDate(0, 0, -7)

		s.True(now.After(sevenDaysEarlier) && now.Before(parsedNextPayment))
	}

}
