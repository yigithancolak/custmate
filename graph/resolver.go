package graph

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *sqlx.DB
}

func NewResolver(db *sqlx.DB) *Resolver {
	return &Resolver{
		DB: db,
	}
}

func (r *mutationResolver) CreateOrganizationResolver(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
	org := &model.Organization{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Email: input.Email,
	}

	// Insert into the database
	query := `INSERT INTO organizations (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.ExecContext(ctx, query, org.ID, org.Name, org.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return org, nil
}
