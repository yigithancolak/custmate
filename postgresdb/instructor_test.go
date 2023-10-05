package postgresdb

import (
	"log"

	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func (s *StoreTestSuite) createRandomInstructor() *model.Instructor {
	organization := s.createRandomOrganization()
	args := model.CreateInstructorInput{
		Name: util.RandomName(),
	}

	instructor, err := s.store.CreateInstructor(organization.ID, args)
	s.NoError(err)
	s.NotEmpty(instructor)
	s.NotNil(instructor)
	s.Equal(args.Name, instructor.Name)

	return instructor
}

func (s *StoreTestSuite) TestCreateInstructor() {
	s.createRandomInstructor()
}

func (s *StoreTestSuite) TestUpdateInstructor() {
	instructor := s.createRandomInstructor()
	name := util.RandomName()
	nextPayment := util.RandomDate()

	args := model.UpdateInstructorInput{
		Name: &name,
	}

	updatedInstructor, err := s.store.UpdateInstructor(instructor.ID, args)
	s.NoError(err)
	s.NotNil(updatedInstructor)
	s.NotEmpty(updatedInstructor)
	s.Equal(name, updatedInstructor.Name)

	log.Println(nextPayment)
}

func (s *StoreTestSuite) TestDeleteInstructor() {
	instructor := s.createRandomInstructor()

	ok, err := s.store.DeleteInstructor(instructor.ID)
	s.NoError(err)
	s.True(ok)

	deletedInstructor, err := s.store.GetInstructorByID(instructor.ID)
	s.Nil(deletedInstructor)
	s.Error(err)
}
