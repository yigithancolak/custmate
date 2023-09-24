package postgresdb

import "github.com/yigithancolak/custmate/graph/model"

func (s *Store) CreateGroupWithTimeTx(group *model.Group, times []*model.Time) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	if err = s.Groups.CreateGroup(group); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	for _, time := range times {
		if err = s.Time.CreateTime(time); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return rbErr
			}
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return nil
}
