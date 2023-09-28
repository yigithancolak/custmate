package postgresdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type TimeStore struct {
	DB *sqlx.DB
}
type queryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

func NewTimeStore(db *sqlx.DB) *TimeStore {
	return &TimeStore{
		DB: db,
	}
}

func (s *TimeStore) CreateTime(q queryer, groupId string, input *model.CreateTimeInput) (*model.Time, error) {
	query := `INSERT INTO times (id, org_group_id, day, start_hour, finish_hour) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	timeId := uuid.New().String() // Assuming you use UUID for your 'id' column

	time := &model.Time{}

	err := q.QueryRow(query, timeId, groupId, input.Day, input.StartHour, input.FinishHour).Scan(&time.ID, &time.GroupID, &time.Day, &time.StartHour, &time.FinishHour)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateTime, err)
	}
	return time, nil
}

func (s *TimeStore) UpdateTime(tx *sql.Tx, time *model.UpdateTimeInput) (*model.Time, error) {
	baseQuery := "UPDATE times SET "
	returnQuery := " RETURNING *"
	var updates []string
	var args []interface{}

	idx := 1

	if time.Day != nil {
		updates = append(updates, fmt.Sprintf("day = $%d", idx))
		args = append(args, time.Day)
		idx++
	}

	if time.StartHour != nil {
		updates = append(updates, fmt.Sprintf("start_hour = $%d", idx))
		args = append(args, time.StartHour)
		idx++
	}

	if time.FinishHour != nil {
		updates = append(updates, fmt.Sprintf("finish_hour = $%d", idx))
		args = append(args, time.FinishHour)
		idx++
	}

	args = append(args, time.ID)
	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery

	var updatedTime model.Time
	if err := tx.QueryRow(query, args...).Scan(&updatedTime.ID, &updatedTime.GroupID, &updatedTime.Day, &updatedTime.StartHour, &updatedTime.FinishHour); err != nil {
		return nil, err
	}

	return &updatedTime, nil
}

func (s *TimeStore) DeleteTime(id string) error {
	query := "DELETE FROM times WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	return err
}
