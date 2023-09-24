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

//DELETE THE **Resolver part of func name before generate

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

func (r *mutationResolver) CreateGroupResolver(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {

	group, err := r.Store.CreateGroupWithTimeTx(input)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *mutationResolver) UpdateGroupResolver(ctx context.Context, id string, input model.UpdateGroupInput) (*model.Group, error) {
	group, err := r.Store.UpdateGroupWithTimeTx(id, input)
	if err != nil {
		return nil, err
	}

	return group, nil

}
