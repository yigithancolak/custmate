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
// 	tokenResponse, err := r.Store.LoginOrganization(email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tokenResponse, nil
// }

// // CreateOrganization is the resolver for the createOrganization field.
// func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
// 	org, err := r.Store.CreateOrganization(input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return org, nil
// }

// // UpdateOrganization is the resolver for the updateOrganization field.
// func (r *mutationResolver) UpdateOrganization(ctx context.Context, input model.UpdateOrganizationInput) (string, error) {
// 	org := middleware.ForContext(ctx)

// 	_, err := r.Store.UpdateOrganization(org.ID, input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageOrganizationUpdated, nil
// }

// // DeleteOrganization is the resolver for the deleteOrganization field.
// func (r *mutationResolver) DeleteOrganization(ctx context.Context) (bool, error) {
// 	org := middleware.ForContext(ctx)

// 	err := r.Store.DeleteOrganization(org.ID)
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
// 	ok, err := r.Store.DeleteGroup(id)

// 	return ok, err
// }

// // CreateInstructor is the resolver for the createInstructor field.
// func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*model.Instructor, error) {
// 	org := middleware.ForContext(ctx)
// 	instructor, err := r.Store.CreateInstructor(org.ID, input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return instructor, err
// }

// // UpdateInstructor is the resolver for the updateInstructor field.
// func (r *mutationResolver) UpdateInstructor(ctx context.Context, id string, input model.UpdateInstructorInput) (string, error) {
// 	_, err := r.Store.UpdateInstructor(id, input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageInstructorUpdated, err
// }

// // DeleteInstructor is the resolver for the deleteInstructor field.
// func (r *mutationResolver) DeleteInstructor(ctx context.Context, id string) (bool, error) {
// 	ok, err := r.Store.DeleteInstructor(id)

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
// 	_, err := r.Store.UpdateCustomerWithTx(id, &input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messageCustomerUpdated, nil
// }

// // DeleteCustomer is the resolver for the deleteCustomer field.
// func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (bool, error) {
// 	err := r.Store.DeleteCustomer(id)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// // CreatePayment is the resolver for the createPayment field.
// func (r *mutationResolver) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (*model.Payment, error) {
// 	org := middleware.ForContext(ctx)

// 	payment, err := r.Store.CreatePaymentTx(org.ID, &input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment, nil
// }

// // UpdatePayment is the resolver for the updatePayment field.
// func (r *mutationResolver) UpdatePayment(ctx context.Context, id string, input model.UpdatePaymentInput) (string, error) {
// 	_, err := r.Store.UpdatePayment(id, &input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}

// 	return messagePaymentUpdated, nil
// }

// // DeletePayment is the resolver for the deletePayment field.
// func (r *mutationResolver) DeletePayment(ctx context.Context, id string) (bool, error) {
// 	err := r.Store.DeletePayment(id)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// // CreateTime is the resolver for the createTime field.
// func (r *mutationResolver) CreateTime(ctx context.Context, groupID string, input model.CreateTimeInput) (*model.Time, error) {
// 	time, err := r.Store.CreateTime(r.Store.DB, groupID, &input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return time, nil
// }

// // UpdateTime is the resolver for the updateTime field.
// func (r *mutationResolver) UpdateTime(ctx context.Context, input model.UpdateTimeInput) (string, error) {
// 	_, err := r.Store.UpdateTime(r.Store.DB, &input)
// 	if err != nil {
// 		return messageUpdateFailed, err
// 	}
// 	return messageTimeUpdated, err
// }

// // DeleteTime is the resolver for the deleteTime field.
// func (r *mutationResolver) DeleteTime(ctx context.Context, id string) (bool, error) {
// 	err := r.Store.DeleteTime(id)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// // GetOrganization is the resolver for the getOrganization field.
// func (r *queryResolver) GetOrganization(ctx context.Context) (*model.Organization, error) {
// 	org := middleware.ForContext(ctx)

// 	return org, nil
// }

// // GetGroup is the resolver for the getGroup field.
// func (r *queryResolver) GetGroup(ctx context.Context, id string) (*model.Group, error) {
// 	group, err := r.Store.GetGroupByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fields := graphql.CollectAllFields(ctx)
// 	if util.Contains[string](fields, "times") {
// 		time, err := r.Store.GetTimesByGroupID(id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		group.Times = time
// 	}

// 	if util.Contains[string](fields, "instructor") {
// 		instructor, err := r.Store.GetInstructorByGroupID(id)
// 		if err != nil {
// 			return nil, err
// 		}

// 		group.Instructor = instructor
// 	}

// 	return group, err
// }

// // ListGroupsByOrganization is the resolver for the listGroupsByOrganization field.
// func (r *queryResolver) ListGroupsByOrganization(ctx context.Context, offset *int, limit *int) (*model.ListGroupsResponse, error) {
// 	org := middleware.ForContext(ctx)
// 	includeTimes := false
// 	includeInstructor := false
// 	includeCustomers := false

// 	fields := GetPreloads(ctx)
// 	if util.Contains[string](fields, "items.times") {
// 		includeTimes = true
// 	}
// 	if util.Contains[string](fields, "items.instructor") {
// 		includeInstructor = true
// 	}
// 	if util.Contains[string](fields, "items.customers") {
// 		includeCustomers = true
// 	}

// 	groups, count, err := r.Store.ListGroupsByFieldID("organization_id", org.ID, offset, limit, includeTimes, includeInstructor, includeCustomers)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListGroupsResponse{
// 		Items:      groups,
// 		TotalCount: count,
// 	}, nil
// }

// // GetInstructor is the resolver for the getInstructor field.
// func (r *queryResolver) GetInstructor(ctx context.Context, id string) (*model.Instructor, error) {
// 	fields := graphql.CollectAllFields(ctx)

// 	instructor, err := r.Store.GetInstructorByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if util.Contains[string](fields, "groups") {
// 		instructor.Groups, err = r.Store.ListGroupsByInstructorID(id)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return instructor, err
// }

// // ListInstructors is the resolver for the listInstructors field.
// func (r *queryResolver) ListInstructors(ctx context.Context, offset *int, limit *int) (*model.ListInstructorsResponse, error) {
// 	org := middleware.ForContext(ctx)
// 	includeGroups := false

// 	fields := GetPreloads(ctx) // Use the same method to get fields
// 	if util.Contains[string](fields, "items.groups") {
// 		includeGroups = true
// 	}

// 	instructors, count, err := r.Store.ListInstructorsByOrganizationID(org.ID, offset, limit, includeGroups)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListInstructorsResponse{
// 		Items:      instructors,
// 		TotalCount: count,
// 	}, nil
// }

// // GetCustomer is the resolver for the getCustomer field.
// func (r *queryResolver) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
// 	includeGroups := false
// 	fields := graphql.CollectAllFields(ctx)
// 	if util.Contains[string](fields, "groups") {
// 		includeGroups = true
// 	}
// 	customer, err := r.Store.GetCustomerByID(id, includeGroups)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return customer, nil
// }

// // ListCustomersByGroup is the resolver for the listCustomersByGroup field.
// func (r *queryResolver) ListCustomersByGroup(ctx context.Context, groupID string, offset *int, limit *int) ([]*model.Customer, error) {
// 	customers, err := r.Store.ListCustomersByGroupID(groupID, offset, limit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return customers, nil
// }

// // ListCustomersByOrganization is the resolver for the listCustomersByOrganization field.
// func (r *queryResolver) ListCustomersByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Customer, error) {
// 	org := middleware.ForContext(ctx)

// 	customers, err := r.Store.ListCustomersByOrganizationID(org.ID, offset, limit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return customers, nil
// }

// // SearchCustomers is the resolver for the searchCustomers field.
// func (r *queryResolver) SearchCustomers(ctx context.Context, filter model.SearchCustomerFilter, offset *int, limit *int) (*model.ListCustomersResponse, error) {
// 	org := middleware.ForContext(ctx)

// 	customers, totalCount, err := r.Store.ListCustomersWithSearchFilter(filter, org.ID, offset, limit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListCustomersResponse{
// 		TotalCount: totalCount,
// 		Items:      customers,
// 	}, nil
// }

// // GetPayment is the resolver for the getPayment field.
// func (r *queryResolver) GetPayment(ctx context.Context, id string) (*model.Payment, error) {
// 	includeCustomer := false
// 	fields := graphql.CollectAllFields(ctx)
// 	if util.Contains[string](fields, "customer") {
// 		includeCustomer = true
// 	}
// 	payment, err := r.Store.GetPaymentByID(id, includeCustomer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment, nil
// }

// // ListPaymentsByOrganization is the resolver for the listPaymentsByOrganization field.
// func (r *queryResolver) ListPaymentsByOrganization(ctx context.Context, offset *int, limit *int, startDate string, endDate string) (*model.ListPaymentsResponse, error) {
// 	org := middleware.ForContext(ctx)
// 	payments, count, err := r.Store.ListPaymentsByFieldID("organization_id", org.ID, offset, limit, startDate, endDate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListPaymentsResponse{
// 		Items:      payments,
// 		TotalCount: count,
// 	}, nil
// }

// // ListPaymentsByGroup is the resolver for the listPaymentsByGroup field.
// func (r *queryResolver) ListPaymentsByGroup(ctx context.Context, groupID string, offset *int, limit *int, startDate string, endDate string) (*model.ListPaymentsResponse, error) {
// 	payments, count, err := r.Store.ListPaymentsByFieldID("org_group_id", groupID, offset, limit, startDate, endDate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListPaymentsResponse{
// 		Items:      payments,
// 		TotalCount: count,
// 	}, nil
// }

// // ListPaymentsByCustomer is the resolver for the listPaymentsByCustomer field.
// func (r *queryResolver) ListPaymentsByCustomer(ctx context.Context, customerID string, offset *int, limit *int, startDate string, endDate string) (*model.ListPaymentsResponse, error) {
// 	payments, count, err := r.Store.ListPaymentsByFieldID("customer_id", customerID, offset, limit, startDate, endDate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListPaymentsResponse{
// 		Items:      payments,
// 		TotalCount: count,
// 	}, nil
// }

// // ListEarningsByOrganization is the resolver for the listEarningsByOrganization field.
// func (r *queryResolver) ListEarningsByOrganization(ctx context.Context, offset *int, limit *int, startDate string, endDate string) (*model.ListEarningsResponse, error) {
// 	org := middleware.ForContext(ctx)
// 	earnings, count, err := r.Store.ListEarningsByOrganization(org.ID, offset, limit, startDate, endDate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.ListEarningsResponse{
// 		Items:      earnings,
// 		TotalCount: count,
// 	}, nil
// }

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
func (r *queryResolver) ListInstructorsByOrganization(ctx context.Context, offset *int, limit *int) ([]*model.Instructor, error) {
	panic(fmt.Errorf("not implemented: ListInstructorsByOrganization - listInstructorsByOrganization"))
}
