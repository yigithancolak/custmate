package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/yigithancolak/custmate/graph/model"
)

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, email string, password string) (*model.TokenResponse, error) {
	panic(fmt.Errorf("not implemented: LoginUser - loginUser"))
}

// CreateOrganization is the resolver for the createOrganization field.
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

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, id string, input model.UpdateOrganizationInput) (*model.Organization, error) {
	resp, err := r.Store.Organizations.UpdateOrganization(id, input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (bool, error) {
	err := r.Store.Organizations.DeleteOrganization(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	group, err := r.Store.CreateGroupWithTimeTx(input)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, input model.UpdateGroupInput) (*model.Group, error) {
	group, err := r.Store.UpdateGroupWithTimeTx(id, input)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteGroup - deleteGroup"))
}

// CreateInstructor is the resolver for the createInstructor field.
func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: CreateInstructor - createInstructor"))
}

// UpdateInstructor is the resolver for the updateInstructor field.
func (r *mutationResolver) UpdateInstructor(ctx context.Context, id string, input model.UpdateInstructorInput) (*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: UpdateInstructor - updateInstructor"))
}

// DeleteInstructor is the resolver for the deleteInstructor field.
func (r *mutationResolver) DeleteInstructor(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteInstructor - deleteInstructor"))
}

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: CreateCustomer - createCustomer"))
}

// UpdateCustomer is the resolver for the updateCustomer field.
func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, input model.UpdateCustomerInput) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: UpdateCustomer - updateCustomer"))
}

// DeleteCustomer is the resolver for the deleteCustomer field.
func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCustomer - deleteCustomer"))
}

// CreatePayment is the resolver for the createPayment field.
func (r *mutationResolver) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented: CreatePayment - createPayment"))
}

// UpdatePayment is the resolver for the updatePayment field.
func (r *mutationResolver) UpdatePayment(ctx context.Context, id string, input model.UpdatePaymentInput) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented: UpdatePayment - updatePayment"))
}

// DeletePayment is the resolver for the deletePayment field.
func (r *mutationResolver) DeletePayment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeletePayment - deletePayment"))
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context, id string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: GetOrganization - getOrganization"))
}

// ListOrganizations is the resolver for the listOrganizations field.
func (r *queryResolver) ListOrganizations(ctx context.Context, offset *int, limit *int) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented: ListOrganizations - listOrganizations"))
}

// GetGroup is the resolver for the getGroup field.
func (r *queryResolver) GetGroup(ctx context.Context, id string) (*model.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroup - getGroup"))
}

// ListGroups is the resolver for the listGroups field.
func (r *queryResolver) ListGroups(ctx context.Context, offset *int, limit *int) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: ListGroups - listGroups"))
}

// ListGroupsByOrganization is the resolver for the listGroupsByOrganization field.
func (r *queryResolver) ListGroupsByOrganization(ctx context.Context, orgID string, offset *int, limit *int) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: ListGroupsByOrganization - listGroupsByOrganization"))
}

// GetInstructor is the resolver for the getInstructor field.
func (r *queryResolver) GetInstructor(ctx context.Context, id string) (*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: GetInstructor - getInstructor"))
}

// ListInstructors is the resolver for the listInstructors field.
func (r *queryResolver) ListInstructors(ctx context.Context, offset *int, limit *int) ([]*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: ListInstructors - listInstructors"))
}

// ListInstructorsByOrganization is the resolver for the listInstructorsByOrganization field.
func (r *queryResolver) ListInstructorsByOrganization(ctx context.Context, orgID string, offset *int, limit *int) ([]*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: ListInstructorsByOrganization - listInstructorsByOrganization"))
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: GetCustomer - getCustomer"))
}

// ListCustomers is the resolver for the listCustomers field.
func (r *queryResolver) ListCustomers(ctx context.Context, offset *int, limit *int) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: ListCustomers - listCustomers"))
}

// ListCustomersByGroup is the resolver for the listCustomersByGroup field.
func (r *queryResolver) ListCustomersByGroup(ctx context.Context, groupID string, offset *int, limit *int) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: ListCustomersByGroup - listCustomersByGroup"))
}

// ListCustomersByOrganization is the resolver for the listCustomersByOrganization field.
func (r *queryResolver) ListCustomersByOrganization(ctx context.Context, orgID string, offset *int, limit *int) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: ListCustomersByOrganization - listCustomersByOrganization"))
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, id string) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented: GetPayment - getPayment"))
}

// ListPaymentsByOrganization is the resolver for the listPaymentsByOrganization field.
func (r *queryResolver) ListPaymentsByOrganization(ctx context.Context, organizationID string, offset *int, limit *int) ([]*model.Payment, error) {
	panic(fmt.Errorf("not implemented: ListPaymentsByOrganization - listPaymentsByOrganization"))
}

// ListPaymentsByGroup is the resolver for the listPaymentsByGroup field.
func (r *queryResolver) ListPaymentsByGroup(ctx context.Context, groupID string, offset *int, limit *int) ([]*model.Payment, error) {
	panic(fmt.Errorf("not implemented: ListPaymentsByGroup - listPaymentsByGroup"))
}

// ListPaymentsByCustomer is the resolver for the listPaymentsByCustomer field.
func (r *queryResolver) ListPaymentsByCustomer(ctx context.Context, customerID string, offset *int, limit *int) ([]*model.Payment, error) {
	panic(fmt.Errorf("not implemented: ListPaymentsByCustomer - listPaymentsByCustomer"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
