package postgresdb

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateInstructor(orgID string, input model.CreateInstructorInput) (*model.Instructor, error) {
	query := `INSERT INTO instructors (id, name, organization_id) VALUES ($1, $2, $3) RETURNING id, name, organization_id`

	instructorId := uuid.New().String()

	var instructor model.Instructor
	err := s.DB.QueryRow(query, instructorId, input.Name, orgID).Scan(&instructor.ID, &instructor.Name, &instructor.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &instructor, nil
}

func (s *Store) UpdateInstructor(id string, input model.UpdateInstructorInput) (*model.Instructor, error) {
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

func (s *Store) DeleteInstructor(id string) (bool, error) {
	query := "DELETE FROM instructors WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Store) GetInstructorByID(id string) (*model.Instructor, error) {
	query := "SELECT id,name, organization_id FROM instructors WHERE id = $1"
	var instructor model.Instructor
	err := s.DB.QueryRow(query, id).Scan(&instructor.ID, &instructor.Name, &instructor.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &instructor, nil
}

func (s *Store) GetInstructorByGroupID(groupID string) (*model.Instructor, error) {
	query := "SELECT id, name, organization_id FROM instructors WHERE id = (SELECT instructor_id FROM org_groups WHERE id = $1)"
	var instructor model.Instructor

	err := s.DB.QueryRow(query, groupID).Scan(&instructor.ID, &instructor.Name, &instructor.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &instructor, nil
}

func (s *Store) ListInstructorsByOrganizationID(orgID string, offset *int, limit *int, includeGroups bool) ([]*model.Instructor, int, error) {
	// Base query without LIMIT and OFFSET
	query := `
	SELECT id, name, COUNT(*) OVER() as total_count
	FROM instructors 
	WHERE organization_id = $1`

	var args []interface{}
	args = append(args, orgID)

	// If limit is provided, append it to the base query and add to the arguments
	if limit != nil {
		query += ` LIMIT $2`
		args = append(args, limit)

		// Only consider offset if limit is provided
		if offset != nil {
			query += ` OFFSET $3`
			args = append(args, offset)
		}
	}

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var instructors []*model.Instructor
	var totalCount int
	for rows.Next() {
		var instructor model.Instructor
		if err := rows.Scan(&instructor.ID, &instructor.Name, &totalCount); err != nil {
			return nil, 0, err
		}

		if includeGroups {
			instructor.Groups, err = s.ListGroupsByInstructorID(instructor.ID)
			if err != nil {
				return nil, 0, err
			}
		}

		instructors = append(instructors, &instructor)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return instructors, totalCount, nil
}
