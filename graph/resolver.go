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

func (r *mutationResolver) CreateGroupResolver(ctx context.Context, input *model.CreateGroupInput) (*model.Group, error) {

	group := &model.Group{
		ID:           uuid.New().String(),
		Name:         input.Name,
		Organization: &model.Organization{ID: input.Organization},
		Instructor:   &model.Instructor{ID: input.Instructor},
	}

	times := make([]*model.Time, len(input.Times))

	for i, t := range input.Times {
		times[i] = &model.Time{
			ID:         uuid.New().String(),
			Day:        t.Day,
			GroupID:    group.ID,
			StartHour:  t.StartHour,
			FinishHour: t.FinishHour,
		}
	}

	group.Times = times

	tx, err := r.Store.DB.Begin()
	if err != nil {
		return nil, err
	}

	if err = r.Store.Groups.CreateGroup(group); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	for _, time := range times {
		if err = r.Store.Time.CreateTime(time); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return nil, rbErr
			}
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, rbErr
		}
		return nil, err
	}

	return group, nil
}
