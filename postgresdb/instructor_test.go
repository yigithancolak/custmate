package postgresdb

import (
	"log"

	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type InstructorTestSuite struct {
	StoreTestSuite
	instructor *model.Instructor
}

func (s *InstructorTestSuite) SetupTest() {
	organization := s.createRandomOrganization()
	s.instructor = s.createRandomInstructor(organization.ID)
}

func (s *StoreTestSuite) createRandomInstructor(organizationID string) *model.Instructor {
	args := model.CreateInstructorInput{
		Name: util.RandomName(),
	}

	instructor, err := s.store.CreateInstructor(organizationID, args)
	s.NoError(err)
	s.NotEmpty(instructor)
	s.NotNil(instructor)
	s.Equal(args.Name, instructor.Name)

	return instructor
}

func (s *InstructorTestSuite) createRandomInstructorsForOneOrganization(n int) []*model.Instructor {
	organization := s.createRandomOrganization()

	var instructors []*model.Instructor
	for i := 0; i < n; i++ {
		createdInstructor := s.createRandomInstructor(organization.ID)

		instructors = append(instructors, createdInstructor)
	}

	return instructors
}

func (s *InstructorTestSuite) TestCreateInstructor() {
	s.NotNil(s.instructor)
	s.NotEmpty(s.instructor)
}

func (s *InstructorTestSuite) TestUpdateInstructor() {
	instructor := s.instructor
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

func (s *InstructorTestSuite) TestDeleteInstructor() {
	instructor := s.instructor

	ok, err := s.store.DeleteInstructor(instructor.ID)
	s.NoError(err)
	s.True(ok)

	deletedInstructor, err := s.store.GetInstructorByID(instructor.ID)
	s.Nil(deletedInstructor)
	s.Error(err)
}

func (s *InstructorTestSuite) TestGetInstructorByID() {
	instructor := s.instructor

	foundInstructor, err := s.store.GetInstructorByID(instructor.ID)
	s.NoError(err)
	s.NotNil(foundInstructor)
	s.NotEmpty(foundInstructor)

	s.Equal(instructor.ID, foundInstructor.ID)
	s.Equal(instructor.Name, foundInstructor.ID)
	s.Equal(instructor.OrganizationID, foundInstructor.OrganizationID)
}

func (s *InstructorTestSuite) TestListInstructorsByOrganizationID() {
	instructors := s.createRandomInstructorsForOneOrganization(8)
	offset := 0
	limit := 10
	instructorsOfOrg, err := s.store.ListInstructorsByOrganizationID(instructors[0].OrganizationID, &offset, &limit, false)
	s.NoError(err)
	s.NotEmpty(instructorsOfOrg)
	s.Equal(len(instructors), len(instructorsOfOrg))

	instructorMap := make(map[string]*model.Instructor)
	for _, ins := range instructors {
		instructorMap[ins.ID] = ins
	}

	for _, retrievedInstructor := range instructorsOfOrg {
		expectedInstructor, exists := instructorMap[retrievedInstructor.ID]
		s.True(exists, "Instructor with ID %s not found in original list", retrievedInstructor.ID)

		s.Equal(expectedInstructor.ID, retrievedInstructor.ID)
		s.Equal(expectedInstructor.Name, retrievedInstructor.Name)
		s.Equal(expectedInstructor.OrganizationID, retrievedInstructor.OrganizationID)
	}

}
