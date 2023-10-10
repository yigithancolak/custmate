package postgresdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
)

func (s *Store) CreateGroup(tx *sql.Tx, input *model.CreateGroupInput, organizationID string) (*model.Group, error) {
	query := `INSERT INTO org_groups (id, name, organization_id, instructor_id) VALUES ($1, $2, $3, $4) RETURNING id, name`

	groupId := uuid.New().String()

	var group model.Group
	err := tx.QueryRow(query, groupId, input.Name, organizationID, input.Instructor).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to insert and scan group: %w", err)
	}

	return &group, nil
}

func (s *Store) UpdateGroup(tx *sql.Tx, id string, input *model.UpdateGroupInput) (*model.Group, error) {
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
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, ErrRollbackTransaction
		}
		return nil, err
	}

	return &group, nil

}

func (s *Store) DeleteGroup(id string) (bool, error) {
	query := "DELETE FROM org_groups WHERE id = $1"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Store) GetGroupByID(id string) (*model.Group, error) {
	query := "SELECT id, name FROM org_groups WHERE id = $1"
	var group model.Group
	err := s.DB.QueryRow(query, id).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

//

func (s *Store) ListGroupsByFieldID(field string, ID string, offset *int, limit *int, includeTimes bool, includeInstructor bool, includeCustomers bool) ([]*model.Group, error) {
	query := fmt.Sprintf("SELECT id, name, instructor_id FROM org_groups WHERE %s = $1 LIMIT $2 OFFSET $3", field)

	rows, err := s.DB.Query(query, ID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*model.Group
	for rows.Next() {
		var group model.Group
		var insID string
		if err := rows.Scan(&group.ID, &group.Name, &insID); err != nil {
			return nil, err
		}

		if includeTimes {
			group.Times, err = s.GetTimesByGroupID(group.ID)
			if err != nil {
				return nil, err
			}
		}

		if includeInstructor {
			group.Instructor, err = s.GetInstructorByID(insID)
			if err != nil {
				return nil, err
			}
		}

		if includeCustomers {
			group.Customers, err = s.GetCustomersByGroupID(group.ID)
			if err != nil {
				return nil, err
			}
		}

		groups = append(groups, &group)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

//

// func (s *Store) ListGroupsByOrganization(organizationID string, offset *int, limit *int, includeTimes bool, includeInstructor bool, includeCustomers bool) ([]*model.Group, error) {
// 	query := "SELECT id, name, instructor_id FROM org_groups WHERE organization_id = $1 LIMIT $2 OFFSET $3"

// 	rows, err := s.DB.Query(query, organizationID, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var groups []*model.Group
// 	for rows.Next() {
// 		var group model.Group
// 		var insID string
// 		if err := rows.Scan(&group.ID, &group.Name, &insID); err != nil {
// 			return nil, err
// 		}

// 		if includeTimes {
// 			group.Times, err = s.GetTimesByGroupID(group.ID)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		if includeInstructor {
// 			group.Instructor, err = s.GetInstructorByID(insID)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		if includeCustomers {
// 			group.Customers, err = s.GetCustomersByGroupID(group.ID)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		groups = append(groups, &group)

// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return groups, nil
// }

func (s *Store) ListGroupsByInstructorID(instructorID string) ([]*model.Group, error) {
	query := "SELECT id, name FROM org_groups WHERE instructor_id = $1"

	rows, err := s.DB.Query(query, instructorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*model.Group
	for rows.Next() {
		var group model.Group
		err = rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		//TODO: INCLUDE TIMES
		groups = append(groups, &group)
	}

	return groups, nil
}

func (s *Store) ListGroupsByCustomerID(customerID string) ([]*model.Group, error) {
	query := ` SELECT og.id, og.name FROM customers c
	JOIN customer_groups cg ON c.id = cg.customer_id
	JOIN org_groups og ON cg.org_group_id = og.id
	WHERE cg.customer_id = $1`

	rows, err := s.DB.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*model.Group

	for rows.Next() {
		var group model.Group
		err := rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		group.Times, err = s.GetTimesByGroupID(group.ID)
		if err != nil {
			return nil, err
		}

		groups = append(groups, &group)
	}

	return groups, nil
}
