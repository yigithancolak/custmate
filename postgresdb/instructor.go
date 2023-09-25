package postgresdb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

type InstructorStore struct {
	DB *sqlx.DB
}

func NewInstructorStore(db *sqlx.DB) *InstructorStore {
	return &InstructorStore{
		DB: db,
	}
}

func (s *InstructorStore) CreateInstructor(input model.CreateInstructorInput) (*model.Instructor, error) {
	query := `INSERT INTO instructors (id, name, organization_id) VALUES ($1, $2, $3) RETURNING id, name, organization_id`

	instructorId := uuid.New().String()

	var instructor model.Instructor
	err := s.DB.QueryRow(query, instructorId, input.Name, input.Organization).Scan(&instructor.ID, &instructor.Name, &instructor.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &instructor, nil
}

func (s *InstructorStore) UpdateInstructor(id string, input model.UpdateInstructorInput) (*model.Instructor, error) {
	baseQuery := "UPDATE instructors SET "
	returnQuery := " RETURNING id, name, organization_id"

	var updates []string
	var args []interface{}

	idx := 1
	if input.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", idx))
		args = append(args, input.Name)
		idx++
	}

	args = append(args, id)
	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery
	var updatedInstructor model.Instructor

	err := s.DB.QueryRow(query, args...).Scan(&updatedInstructor.ID, &updatedInstructor.Name, &updatedInstructor.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &updatedInstructor, nil
}

func (s *InstructorStore) DeleteInstructor(id string) (bool, error) {
	query := "DELETE FROM instructors WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
