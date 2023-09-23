package postgresdb

import "github.com/jmoiron/sqlx"

type Store struct {
	DB            *sqlx.DB
	Organizations *OrganizationStore
	Groups        *GroupStore
	Time          *TimeStore
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		DB:            db,
		Organizations: NewOrganizationStore(db),
		Groups:        NewGroupStore(db),
		Time:          NewTimeStore(db),
	}
}
