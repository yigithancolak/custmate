package postgresdb

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func createRandomInstructor(t *testing.T) *model.Instructor {
	requires := require.New(t)
	organization := createRandomOrganization(t)
	args := model.CreateInstructorInput{
		Name: util.RandomName(),
	}

	instructor, err := testStore.CreateInstructor(organization.ID, args)
	requires.NoError(err)
	requires.NotEmpty(instructor)
	requires.NotNil(instructor)
	requires.Equal(args.Name, instructor.Name)

	return instructor
}

func TestCreateInstructor(t *testing.T) {
	createRandomInstructor(t)
}

func TestUpdateInstructor(t *testing.T) {
	requires := require.New(t)
	instructor := createRandomInstructor(t)
	name := util.RandomName()
	nextPayment := util.RandomDate()

	args := model.UpdateInstructorInput{
		Name: &name,
	}

	updatedInstructor, err := testStore.UpdateInstructor(instructor.ID, args)
	requires.NoError(err)
	requires.NotNil(updatedInstructor)
	requires.NotEmpty(updatedInstructor)
	requires.Equal(name, updatedInstructor.Name)

	log.Println(nextPayment)
}

func TestDeleteInstructor(t *testing.T) {
	requires := require.New(t)
	instructor := createRandomInstructor(t)

	ok, err := testStore.DeleteInstructor(instructor.ID)
	requires.NoError(err)
	requires.True(ok)

	deletedInstructor, err := testStore.GetInstructorByID(instructor.ID)
	requires.Nil(deletedInstructor)
	requires.Error(err)
}
