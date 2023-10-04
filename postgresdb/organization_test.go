package postgresdb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func createRandomOrganization(t *testing.T) *model.Organization {
	requires := require.New(t)

	args := model.CreateOrganizationInput{
		Name:     util.RandomName(),
		Email:    util.RandomMail(),
		Password: util.RandomPassword(),
	}

	account, err := testStore.CreateOrganization(args)
	requires.NoError(err)
	requires.NotEmpty(account)

	requires.Equal(args.Name, account.Name)
	requires.Equal(args.Email, account.Email)

	return account
}

func TestCreateOrganization(t *testing.T) {
	createRandomOrganization(t)
}

func TestUpdateOrganization(t *testing.T) {
	requires := require.New(t)
	organization := createRandomOrganization(t)
	name := util.RandomName()
	email := util.RandomMail()
	password := util.RandomPassword()

	args := model.UpdateOrganizationInput{
		Name:     &name,
		Email:    &email,
		Password: &password,
	}

	updatedOrg, err := testStore.UpdateOrganization(organization.ID, args)
	requires.NoError(err)
	requires.NotEmpty(updatedOrg)

	requires.Equal(name, updatedOrg.Name)
	requires.Equal(email, updatedOrg.Email)
	requires.Equal(organization.ID, updatedOrg.ID)
}

func TestDeleteOrganization(t *testing.T) {
	requires := require.New(t)
	organization := createRandomOrganization(t)

	err := testStore.DeleteOrganization(organization.ID)
	requires.NoError(err)

	deletedOrg, err := testStore.GetOrganizationById(organization.ID)
	requires.Nil(deletedOrg)
	requires.Error(err)
}
