package postgresdb

import "github.com/jmoiron/sqlx"

type Store struct {
	Organizations *OrganizationStore
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Organizations: NewOrganizationStore(db),
	}
}
