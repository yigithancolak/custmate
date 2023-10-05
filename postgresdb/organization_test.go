package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func (s *StoreTestSuite) createRandomOrganization() *model.Organization {
	args := model.CreateOrganizationInput{
		Name:     util.RandomName(),
		Email:    util.RandomMail(),
		Password: util.RandomPassword(),
	}

	account, err := s.store.CreateOrganization(args)
	s.NoError(err)
	s.NotEmpty(account)

	s.Equal(args.Name, account.Name)
	s.Equal(args.Email, account.Email)

	return account
}

func (s *StoreTestSuite) TestCreateOrganization() {
	s.createRandomOrganization()
}

func (s *StoreTestSuite) TestUpdateOrganization() {
	organization := s.createRandomOrganization()
	name := util.RandomName()
	email := util.RandomMail()
	password := util.RandomPassword()

	args := model.UpdateOrganizationInput{
		Name:     &name,
		Email:    &email,
		Password: &password,
	}

	updatedOrg, err := s.store.UpdateOrganization(organization.ID, args)
	s.NoError(err)
	s.NotEmpty(updatedOrg)

	s.Equal(name, updatedOrg.Name)
	s.Equal(email, updatedOrg.Email)
	s.Equal(organization.ID, updatedOrg.ID)
}

func (s *StoreTestSuite) TestDeleteOrganization() {
	organization := s.createRandomOrganization()

	err := s.store.DeleteOrganization(organization.ID)
	s.NoError(err)

	deletedOrg, err := s.store.GetOrganizationById(organization.ID)
	s.Nil(deletedOrg)
	s.Error(err)
}
