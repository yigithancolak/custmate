package graph

//go:generate go run github.com/99designs/gqlgen

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

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.TokenResponse, error) {
	tokenResponse, err := r.Store.Organizations.LoginOrganization(email, password)
	if err != nil {
		return nil, err
	}
	return tokenResponse, nil
}

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
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

func (r *mutationResolver) UpdateOrganization(ctx context.Context, id string, input model.UpdateOrganizationInput) (*model.Organization, error) {
	resp, err := r.Store.Organizations.UpdateOrganization(id, input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (bool, error) {
	err := r.Store.Organizations.DeleteOrganization(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {

	group, err := r.Store.CreateGroupWithTimeTx(input)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, input model.UpdateGroupInput) (*model.Group, error) {
	group, err := r.Store.UpdateGroupWithTimeTx(id, input)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (bool, error) {
	ok, err := r.Store.Groups.DeleteGroup(id)

	return ok, err
}

func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*model.Instructor, error) {
	instructor, err := r.Store.Instructors.CreateInstructor(input)
	if err != nil {
		return nil, err
	}

	return instructor, err
}

func (r *mutationResolver) UpdateInstructor(ctx context.Context, id string, input model.UpdateInstructorInput) (*model.Instructor, error) {
	instructor, err := r.Store.Instructors.UpdateInstructor(id, input)
	if err != nil {
		return nil, err
	}

	return instructor, err
}

func (r *mutationResolver) DeleteInstructor(ctx context.Context, id string) (bool, error) {
	ok, err := r.Store.Instructors.DeleteInstructor(id)

	return ok, err
}
