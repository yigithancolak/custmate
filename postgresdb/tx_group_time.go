package postgresdb

import (
	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateGroupWithTimeTx(input model.CreateGroupInput, organizationID string) (*model.Group, error) {

	tx, err := s.DB.Begin()
	if err != nil {
		return nil, ErrBeginTransaction
	}
	createdGroup, err := s.Groups.CreateGroup(tx, &input, organizationID)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	for _, t := range input.Times {
		createdTime, err := s.Time.CreateTime(tx, createdGroup.ID, t)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, ErrRollbackTransaction
			}
			return nil, err
		}
		createdGroup.Times = append(createdGroup.Times, createdTime)
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	return createdGroup, nil
}

func (s *Store) UpdateGroupWithTimeTx(id string, groupInput model.UpdateGroupInput) (*model.Group, error) {

	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	group, err := s.Groups.UpdateGroup(tx, id, &groupInput)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	var times []*model.Time
	for _, t := range groupInput.Times {
		time, err := s.Time.UpdateTime(tx, t)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		times = append(times, time)
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	group.Times = times

	//s.Instructor.GetInstructor

	return group, nil
}