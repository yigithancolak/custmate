package postgresdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type GroupStore struct {
	DB *sqlx.DB
}

func NewGroupStore(db *sqlx.DB) *GroupStore {
	return &GroupStore{
		DB: db,
	}
}

func (s *GroupStore) CreateGroup(tx *sql.Tx, input *model.CreateGroupInput, organizationID string) (*model.Group, error) {
	query := `INSERT INTO org_groups (id, name, organization_id, instructor_id) VALUES ($1, $2, $3, $4) RETURNING id, name`

	groupId := uuid.New().String()

	group := &model.Group{}
	err := tx.QueryRow(query, groupId, input.Name, organizationID, input.Instructor).Scan(&group.ID, &group.Name)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, ErrInsertScanGroup
	}

	return group, nil
}

func (s *GroupStore) UpdateGroup(tx *sql.Tx, id string, input *model.UpdateGroupInput) (*model.Group, error) {
	baseQuery := "UPDATE org_groups SET "
	returnQuery := " RETURNING id,name"
	var updates []string
	var args []interface{}

	idx := 1

	if input.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", idx))
		args = append(args, input.Name)
		idx++
	}

	if input.Instructor != nil {
		updates = append(updates, fmt.Sprintf("instructor_id = $%d", idx))
		args = append(args, input.Instructor)
		idx++
	}

	args = append(args, id)
	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery

	var group model.Group

	err := tx.QueryRow(query, args...).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}

	return &group, nil

}

func (s *GroupStore) DeleteGroup(id string) (bool, error) {
	query := "DELETE FROM org_groups WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
