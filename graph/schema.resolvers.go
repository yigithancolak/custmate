package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/yigithancolak/custmate/graph/model"
)

// // Login is the resolver for the login field.
// func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.TokenResponse, error) {
// 	tokenResponse, err := r.Store.Organizations.LoginOrganization(email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tokenResponse, nil
// }

// // CreateOrganization is the resolver for the createOrganization field.
// func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
// 	org, err := r.Store.Organizations.CreateOrganization(input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return org, nil
// }

// // UpdateOrganization is the resolver for the updateOrganization field.
// func (r *mutationResolver) UpdateOrganization(ctx context.Context, input model.UpdateOrganizationInput) (string, error) {
// 	org := middleware.ForContext(ctx)

// 	_, err := r.Store.Organizations.UpdateOrganization(org.ID, input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageOrganizationUpdated, nil
// }

// // DeleteOrganization is the resolver for the deleteOrganization field.
// func (r *mutationResolver) DeleteOrganization(ctx context.Context) (bool, error) {
// 	org := middleware.ForContext(ctx)

// 	err := r.Store.Organizations.DeleteOrganization(org.ID)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// // CreateGroup is the resolver for the createGroup field.
// func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
// 	org := middleware.ForContext(ctx)

// 	group, err := r.Store.CreateGroupWithTimeTx(input, org.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fields := graphql.CollectAllFields(ctx)

// 	if util.Contains[string](fields, "instructor") {
// 		//TODO: ADD INSTRUCTOR TO RETURN OBJECT
// 		log.Println("instructor wanted")
// 	}

// 	return group, nil
// }

// // UpdateGroup is the resolver for the updateGroup field.
// func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, input model.UpdateGroupInput) (string, error) {
// 	err := r.Store.UpdateGroupWithTimeTx(id, input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageGroupUpdated, nil
// }

// // DeleteGroup is the resolver for the deleteGroup field.
// func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (bool, error) {
// 	ok, err := r.Store.Groups.DeleteGroup(id)

// 	return ok, err
// }

// // CreateInstructor is the resolver for the createInstructor field.
// func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*model.Instructor, error) {
// 	org := middleware.ForContext(ctx)
// 	instructor, err := r.Store.Instructors.CreateInstructor(org.ID, input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return instructor, err
// }

// // UpdateInstructor is the resolver for the updateInstructor field.
// func (r *mutationResolver) UpdateInstructor(ctx context.Context, id string, input model.UpdateInstructorInput) (string, error) {
// 	_, err := r.Store.Instructors.UpdateInstructor(id, input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageInstructorUpdated, err
// }

// // DeleteInstructor is the resolver for the deleteInstructor field.
// func (r *mutationResolver) DeleteInstructor(ctx context.Context, id string) (bool, error) {
// 	ok, err := r.Store.Instructors.DeleteInstructor(id)

// 	return ok, err
// }

// // CreateCustomer is the resolver for the createCustomer field.
// func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
// 	org := middleware.ForContext(ctx)

// 	customer, err := r.Store.CreateCustomerWithTx(&input, org.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return customer, nil
// }

// // UpdateCustomer is the resolver for the updateCustomer field.
// func (r *mutationResolver) UpdateCustomer(ctx context.Context, id string, input model.UpdateCustomerInput) (string, error) {
// 	err := r.Store.UpdateCustomerWithTx(id, &input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageCustomerUpdated, nil
// }

// // DeleteCustomer is the resolver for the deleteCustomer field.
// func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (bool, error) {
// 	err := r.Store.Customers.DeleteCustomer(id)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// // CreatePayment is the resolver for the createPayment field.
// func (r *mutationResolver) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (*model.Payment, error) {
// 	panic(fmt.Errorf("not implemented: CreatePayment - createPayment"))
// }

// UpdatePayment is the resolver for the updatePayment field.
func (r *mutationResolver) UpdatePayment(ctx context.Context, id string, input model.UpdatePaymentInput) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatePayment - updatePayment"))
}

// DeletePayment is the resolver for the deletePayment field.
func (r *mutationResolver) DeletePayment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeletePayment - deletePayment"))
}

// CreateTime is the resolver for the createTime field.
func (r *mutationResolver) CreateTime(ctx context.Context, input model.CreateTimeInput) (*model.Time, error) {
	panic(fmt.Errorf("not implemented: CreateTime - createTime"))
}

// UpdateTime is the resolver for the updateTime field.
func (r *mutationResolver) UpdateTime(ctx context.Context, id string, input model.UpdateTimeInput) (string, error) {
	panic(fmt.Errorf("not implemented: UpdateTime - updateTime"))
}

// DeleteTime is the resolver for the deleteTime field.
func (r *mutationResolver) DeleteTime(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTime - deleteTime"))
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: GetOrganization - getOrganization"))
}

// GetGroup is the resolver for the getGroup field.
func (r *queryResolver) GetGroup(ctx context.Context, id string) (*model.Group, error) {
	panic(fmt.Errorf("not implemented: GetGroup - getGroup"))
}

// ListGroupsByOrganization is the resolver for the listGroupsByOrganization field.
func (r *queryResolver) ListGroupsByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Group, error) {
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
func (r *queryResolver) ListInstructorsByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: ListInstructorsByOrganization - listInstructorsByOrganization"))
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: GetCustomer - getCustomer"))
}

// ListCustomersByGroup is the resolver for the listCustomersByGroup field.
func (r *queryResolver) ListCustomersByGroup(ctx context.Context, groupID string, offset *int, limit *int) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: ListCustomersByGroup - listCustomersByGroup"))
}

// ListCustomersByOrganization is the resolver for the listCustomersByOrganization field.
func (r *queryResolver) ListCustomersByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: ListCustomersByOrganization - listCustomersByOrganization"))
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, id string) (*model.Payment, error) {
	panic(fmt.Errorf("not implemented: GetPayment - getPayment"))
}

// ListPaymentsByOrganization is the resolver for the listPaymentsByOrganization field.
func (r *queryResolver) ListPaymentsByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Payment, error) {
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
