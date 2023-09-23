package graph

import (
	"context"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/postgresdb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store *postgresdb.Store
}

func NewResolver(store *postgresdb.Store) *Resolver {
	return &Resolver{
		Store: store,
	}
}

func (r *mutationResolver) CreateOrganizationResolver(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
	org := &model.Organization{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Email: input.Email,
	}

	// Insert into the database
	err := r.Store.Organizations.CreateOrganization(org, input.Password)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (r *mutationResolver) UpdateOrganizationResolver(ctx context.Context, id string, input model.UpdateOrganizationInput) (*model.Organization, error) {
	resp, err := r.Store.Organizations.UpdateOrganization(id, input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *mutationResolver) DeleteOrganizationResolver(ctx context.Context, id string) (bool, error) {
	err := r.Store.Organizations.DeleteOrganization(id)
	if err != nil {
		return false, err
	}
	return true, nil
}
