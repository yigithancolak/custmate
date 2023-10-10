package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type OrganizationTestSuite struct {
	StoreTestSuite
	organization *model.Organization
}

func (s *OrganizationTestSuite) SetupTest() {
	s.organization = s.createRandomOrganization()
}

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

func (s *OrganizationTestSuite) TestCreateOrganization() {
	s.NotNil(s.organization)
	s.NotEmpty(s.organization)
}

func (s *OrganizationTestSuite) TestUpdateOrganization() {
	organization := s.organization
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

	s.Equal(organization.ID, updatedOrg.ID)
	s.Equal(name, updatedOrg.Name)
	s.Equal(email, updatedOrg.Email)
}

func (s *OrganizationTestSuite) TestDeleteOrganization() {
	organization := s.organization

	err := s.store.DeleteOrganization(organization.ID)
	s.NoError(err)

	deletedOrg, err := s.store.GetOrganizationById(organization.ID)
	s.Nil(deletedOrg)
	s.Error(err)
}

func (s *OrganizationTestSuite) TestGetOrganizationByID() {
	organization := s.organization

	foundOrganization, err := s.store.GetOrganizationById(organization.ID)
	s.NoError(err)
	s.NotNil(foundOrganization)
	s.NotEmpty(foundOrganization)

	s.Equal(organization.ID, foundOrganization.ID)
	s.Equal(organization.Email, foundOrganization.Email)
	s.Equal(organization.Name, foundOrganization.Name)
}
