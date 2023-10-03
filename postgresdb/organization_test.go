package postgresdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func createRandomOrganization(t *testing.T) *model.Organization {
	asserts := assert.New(t)

	args := model.CreateOrganizationInput{
		Name:     util.RandomName(),
		Email:    util.RandomMail(),
		Password: util.RandomPassword(),
	}

	account, err := testStore.CreateOrganization(args)
	asserts.NoError(err)
	asserts.NotEmpty(account)

	asserts.Equal(args.Name, account.Name)
	asserts.Equal(args.Email, account.Email)

	return account
}

func TestCreateOrganization(t *testing.T) {
	createRandomOrganization(t)
}

func TestUpdateOrganization(t *testing.T) {
	asserts := assert.New(t)
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
	asserts.NoError(err)
	asserts.NotEmpty(updatedOrg)

	asserts.Equal(name, updatedOrg.Name)
	asserts.Equal(email, updatedOrg.Email)
	asserts.Equal(organization.ID, updatedOrg.ID)
}
