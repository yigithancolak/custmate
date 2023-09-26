package postgresdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/token"
	"github.com/yigithancolak/custmate/util"
)

type OrganizationStore struct {
	DB       *sqlx.DB
	JWTMaker *token.JWTMaker
}

func NewOrganizationStore(db *sqlx.DB, jwtMaker *token.JWTMaker) *OrganizationStore {
	return &OrganizationStore{
		DB:       db,
		JWTMaker: jwtMaker,
	}
}

func (s *OrganizationStore) LoginOrganization(email, password string) (*model.TokenResponse, error) {
	query := "SELECT id,password FROM organizations WHERE email = $1"

	var orgID string
	var foundPassword string
	err := s.DB.QueryRow(query, email).Scan(&orgID, &foundPassword)
	if err != nil {
		return nil, err
	}

	if err = util.ComparePassword(password, foundPassword); err != nil {
		return nil, errors.New("wrong credentials")
	}

	accessToken, _, err := s.JWTMaker.CreateToken(orgID, s.JWTMaker.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{AccessToken: accessToken}, nil
}

func (s *OrganizationStore) CreateOrganization(input model.CreateOrganizationInput) (*model.Organization, error) {
	query := `INSERT INTO organizations (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id, name, email`
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	orgId := uuid.New().String()
	var org model.Organization
	err = s.DB.QueryRow(query, orgId, input.Name, input.Email, hashedPassword).Scan(&org.ID, &org.Name, &org.Email)
	return &org, nil
}

func (s *OrganizationStore) UpdateOrganization(orgID string, input model.UpdateOrganizationInput) (*model.Organization, error) {
	baseQuery := "UPDATE organizations SET "
	returnQuery := " RETURNING id, name, email"

	// slices to build the dynamic part of the query
	var updates []string
	var args []interface{}

	idx := 1 // Placeholder index counter

	if input.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", idx))
		args = append(args, *input.Name)
		idx++
	}

	if input.Email != nil {
		updates = append(updates, fmt.Sprintf("email = $%d", idx))
		args = append(args, *input.Email)
		idx++
	}

	// Check if any field was provided to update
	if len(updates) == 0 {
		return nil, errors.New("no fields provided to update")
	}

	args = append(args, orgID)
	query := baseQuery + strings.Join(updates, ", ") + fmt.Sprintf(" WHERE id = $%d", idx) + returnQuery

	var org model.Organization
	err := s.DB.QueryRow(query, args...).Scan(&org.ID, &org.Name, &org.Email)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (s *OrganizationStore) DeleteOrganization(id string) error {
	query := `DELETE FROM organizations WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}
