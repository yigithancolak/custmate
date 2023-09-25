package postgresdb

import (
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/token"
)

type Store struct {
	DB            *sqlx.DB
	Organizations *OrganizationStore
	Groups        *GroupStore
	Time          *TimeStore
	Instructors   *InstructorStore
	JWTMaker      *token.JWTMaker
}

func NewStore(db *sqlx.DB, jwtMaker *token.JWTMaker) *Store {
	return &Store{
		DB:            db,
		Organizations: NewOrganizationStore(db),
		Groups:        NewGroupStore(db),
		Time:          NewTimeStore(db),
		Instructors:   NewInstructorStore(db),
		JWTMaker:      jwtMaker,
	}
}
