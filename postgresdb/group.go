package postgresdb

import (
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

func (s *GroupStore) CreateGroup(group *model.Group) error {

	query := `INSERT INTO org_groups (id, name, organization_id, instructor_id) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(query, group.ID, group.Name, group.Organization.ID, group.Instructor.ID)

	return err
}
