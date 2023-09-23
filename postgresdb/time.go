package postgresdb

import (
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type TimeStore struct {
	DB *sqlx.DB
}

func NewTimeStore(db *sqlx.DB) *TimeStore {
	return &TimeStore{
		DB: db,
	}
}

func (s *TimeStore) CreateTime(time *model.Time) error {

	query := `INSERT INTO times (id, org_group_id, day, start_hour, finish_hour) VALUES ($1, $2, $3, $4, $5);`
	_, err := s.DB.Exec(query, time.ID, time.GroupID, time.Day, time.StartHour, time.FinishHour)

	return err
}
