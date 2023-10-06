package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/util"
)

type TimeTestSuite struct {
	StoreTestSuite
	Group *model.Group
}

func (s *TimeTestSuite) SetupTest() {
	organization := s.createRandomOrganization()
	instructor := s.createRandomInstructor(organization.ID)
	s.Group = s.createRandomGroup(instructor.ID, organization.ID)

	time := s.createRandomTime(s.Group.ID)

	s.Group.Times = append(s.Group.Times, time)

}

func (s *TimeTestSuite) createRandomTime(groupID string) *model.Time {
	args := &model.CreateTimeInput{
		Day:        util.RandomDay(),
		StartHour:  util.RandomTime(),
		FinishHour: util.RandomTime(),
	}

	time, err := s.store.CreateTime(s.store.DB, groupID, args)
	s.NoError(err)
	s.NotEmpty(time)
	s.NotNil(time)

	s.Equal(args.Day, time.Day)
	s.Equal(args.StartHour, time.StartHour)
	s.Equal(args.FinishHour, time.FinishHour)

	return time
}

func (s *TimeTestSuite) TestUpdateTime() {
	timeID := s.Group.Times[util.RandomIntBetween(0, int64(len(s.Group.Times)-1))].ID
	day := util.RandomDay()
	startHour := util.RandomTime()
	finishHour := util.RandomTime()

	args := model.UpdateTimeInput{
		ID:         timeID,
		Day:        &day,
		StartHour:  &startHour,
		FinishHour: &finishHour,
	}

	updatedTime, err := s.store.UpdateTime(s.store.DB, &args)
	s.NoError(err)
	s.Equal(args.ID, updatedTime.ID)
	s.Equal(args.Day, updatedTime.Day)
	s.Equal(args.StartHour, updatedTime.StartHour)
	s.Equal(args.FinishHour, updatedTime.FinishHour)
}

func (s *TimeTestSuite) TestGetTimesByGroupID() {
	groupID := s.Group.ID

	foundTimes, err := s.store.GetTimesByGroupID(groupID)
	s.NoError(err)

	timesMap := make(map[string]*model.Time, len(s.Group.Times))

	for _, t := range s.Group.Times {
		timesMap[t.ID] = t
	}

	for _, foundTime := range foundTimes {
		expectedTime, existing := timesMap[foundTime.ID]
		s.True(existing)

		s.Equal(expectedTime.ID, foundTime.ID)
		s.Equal(expectedTime.GroupID, foundTime.GroupID)
		s.Equal(expectedTime.Day, foundTime.Day)
		s.Equal(expectedTime.StartHour, foundTime.StartHour)
		s.Equal(expectedTime.FinishHour, foundTime.FinishHour)
	}

}
