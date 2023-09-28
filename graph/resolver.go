package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/middleware"
	"github.com/yigithancolak/custmate/postgresdb"
	"github.com/yigithancolak/custmate/util"
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

	org, err := r.Store.Organizations.CreateOrganization(input)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (r *mutationResolver) UpdateOrganization(ctx context.Context, input model.UpdateOrganizationInput) (*model.Organization, error) {
	org := middleware.ForContext(ctx)

	resp, err := r.Store.Organizations.UpdateOrganization(org.ID, input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context) (bool, error) {
	org := middleware.ForContext(ctx)

	err := r.Store.Organizations.DeleteOrganization(org.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {

	org := middleware.ForContext(ctx)

	group, err := r.Store.CreateGroupWithTimeTx(input, org.ID)
	if err != nil {
		return nil, err
	}

	fields := graphql.CollectAllFields(ctx)

	if util.Contains[string](fields, "instructor") {
		//TODO: ADD INSTRUCTOR TO RETURN OBJECT
		log.Println("instructor wanted")
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
	org := middleware.ForContext(ctx)
	instructor, err := r.Store.Instructors.CreateInstructor(org.ID, input)
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

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	org := middleware.ForContext(ctx)

	customer, err := r.Store.CreateCustomerWithTx(&input, org.ID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, input model.UpdateCustomerInput) (string, error) {
	err := r.Store.UpdateCustomerWithTx(id, &input)
	if err != nil {
		return "Uptade failed", err
	}

	return "Customer updated", nil
}
