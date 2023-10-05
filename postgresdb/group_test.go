package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type GroupTestSuite struct {
	StoreTestSuite
	Organization *model.Organization
	Instructor   *model.Instructor
	Group        *model.Group
}

func (s *GroupTestSuite) SetupTest() {
	s.Organization = s.createRandomOrganization()
	s.Instructor = s.createRandomInstructor(s.Organization.ID)
	s.Group = s.createRandomGroup(s.Instructor.ID, s.Organization.ID)

}

func (s *StoreTestSuite) createRandomGroup(instructorID, organizationID string) *model.Group {

	args := model.CreateGroupInput{
		Name:       util.RandomName(),
		Instructor: instructorID,
		Times: []*model.CreateTimeInput{
			{
				Day:        util.RandomDay(),
				StartHour:  util.RandomTime(),
				FinishHour: util.RandomTime(),
			},
			{
				Day:        util.RandomDay(),
				StartHour:  util.RandomTime(),
				FinishHour: util.RandomTime(),
			},
		},
	}

	group, err := s.store.CreateGroupWithTimeTx(args, organizationID)
	s.NoError(err)
	s.NotNil(group)

	s.Equal(args.Name, group.Name)
	//TODO: CHECK INSTRUCTOR

	for i, timeOfGroup := range group.Times {
		s.Equal(args.Times[i].Day, timeOfGroup.Day)
		s.Equal(args.Times[i].StartHour+":00", timeOfGroup.StartHour)
		s.Equal(args.Times[i].FinishHour+":00", timeOfGroup.FinishHour)
		//added ":00" because the database time format
	}

	return group
}

func (s *GroupTestSuite) TestCreateGroup() {
	s.NotEmpty(s.Group)
	s.NotNil(s.Group)
}

func (s *GroupTestSuite) TestUpdateGroup() {
	group := s.Group

	name := util.RandomName()
	var times []*model.UpdateTimeInput

	for _, t := range group.Times {
		day := util.RandomDay()
		start := util.RandomTime()
		finish := util.RandomTime()

		timeArgs := model.UpdateTimeInput{
			ID:         t.ID,
			Day:        &day,
			StartHour:  &start,
			FinishHour: &finish,
		}

		times = append(times, &timeArgs)
	}

	newInstructor := s.createRandomInstructor(s.Organization.ID)

	args := model.UpdateGroupInput{
		Name:       &name,
		Instructor: &newInstructor.ID,
		Times:      times,
	}

	err := s.store.UpdateGroupWithTimeTx(group.ID, args)
	s.NoError(err)
	//TODO: MAKE UPDATEGROUPWITHTIMETX TESTABLE WITH RETURNING GROUP
}

func (s *GroupTestSuite) TestDeleteGroup() {
	group := s.Group

	ok, err := s.store.DeleteGroup(group.ID)
	s.NoError(err)
	s.True(ok)
}
