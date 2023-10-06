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

func (s *GroupTestSuite) createRandomGroupsForOrgAndInstructor(n int) []*model.Group {
	organization := s.Organization
	instructor := s.Instructor

	var groups []*model.Group
	for i := 0; i < n; i++ {
		g := s.createRandomGroup(instructor.ID, organization.ID)
		groups = append(groups, g)
	}

	s.Equal(n, len(groups))

	return groups
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

func (s *GroupTestSuite) TestGetGroupByID() {
	group := s.Group

	foundGroup, err := s.store.GetGroupByID(group.ID)
	s.NoError(err)
	s.NotNil(foundGroup)

	s.Equal(group.ID, foundGroup.ID)
	s.Equal(group.Name, foundGroup.Name)
	s.Equal(len(group.Times), len(foundGroup.Times))

	timesMap := make(map[string]*model.Time, len(group.Times))

	for _, t := range group.Times {
		timesMap[t.ID] = t
	}

	for _, t := range foundGroup.Times {
		expectedTime, existing := timesMap[t.ID]
		s.True(existing)

		s.Equal(expectedTime.ID, t.ID)
		s.Equal(expectedTime.Day, t.Day)
		s.Equal(expectedTime.StartHour, t.StartHour)
		s.Equal(expectedTime.FinishHour, t.FinishHour)
	}

}

func (s *GroupTestSuite) TestListGroupsByOrganizationID() {
	n := 10
	groups := s.createRandomGroupsForOrgAndInstructor(n)
	s.NotNil(groups)
	s.NotEmpty(groups)

	groupsMap := make(map[string]*model.Group, len(groups))

	for _, g := range groups {
		groupsMap[g.ID] = g
	}

	offset := 0
	limit := n

	foundGroups, err := s.store.ListGroupsByOrganization(s.Organization.ID, &offset, &limit, false, false, false)
	s.NoError(err)
	s.NotNil(foundGroups)
	s.NotEmpty(foundGroups)

	for _, foundGroup := range foundGroups {
		expectedGroup, existing := groupsMap[foundGroup.ID]
		s.True(existing)

		s.Equal(expectedGroup.ID, foundGroup.ID)
		s.Equal(expectedGroup.Name, foundGroup.Name)
	}

}

func (s *GroupTestSuite) TestListGroupsByInstructorID() {
	n := 10
	groups := s.createRandomGroupsForOrgAndInstructor(n)

	foundGroups, err := s.store.ListGroupsByInstructorID(s.Instructor.ID)
	s.NoError(err)
	s.NotEmpty(foundGroups)
	s.NotNil(foundGroups)

	s.Equal(len(groups), len(foundGroups))

	groupsMap := make(map[string]*model.Group)
	for _, g := range groups {
		groupsMap[g.ID] = g
	}

	for _, foundGroup := range foundGroups {
		expectedGroup, existing := groupsMap[foundGroup.ID]
		s.True(existing)

		s.Equal(expectedGroup.ID, foundGroup.ID)
		s.Equal(expectedGroup.Name, foundGroup.Name)

	}

}
