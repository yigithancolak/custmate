package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

func (s *StoreTestSuite) createRandomGroup() *model.Group {
	instructor := s.createRandomInstructor()
	args := model.CreateGroupInput{
		Name:       util.RandomName(),
		Instructor: instructor.ID,
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

	group, err := s.store.CreateGroupWithTimeTx(args, instructor.OrganizationID)
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

func (s *StoreTestSuite) TestCreateGroup() {
	s.createRandomGroup()
}

func (s *StoreTestSuite) TestUpdateGroup() {

	group := s.createRandomGroup()
	instructor := s.createRandomInstructor() //created for changing the instructor

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

	args := model.UpdateGroupInput{
		Name:       &name,
		Instructor: &instructor.ID,
		Times:      times,
	}

	err := s.store.UpdateGroupWithTimeTx(group.ID, args)
	s.NoError(err)
	//TODO: MAKE UPDATEGROUPWITHTIMETX TESTABLE WITH RETURNING GROUP
}

func (s *StoreTestSuite) TestDeleteGroup() {
	group := s.createRandomGroup()

	ok, err := s.store.DeleteGroup(group.ID)
	s.NoError(err)
	s.True(ok)
}
