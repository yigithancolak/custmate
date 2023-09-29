package postgresdb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/token"
)

type Store struct {
	DB            *sqlx.DB
	Organizations *OrganizationStore
	Groups        *GroupStore
	Time          *TimeStore
	Instructors   *InstructorStore
	Customers     *CustomerStore
	Payments      *PaymentStore
	JWTMaker      *token.JWTMaker
}

type queryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

func NewStore(db *sqlx.DB, jwtMaker *token.JWTMaker) *Store {
	return &Store{
		DB:            db,
		Organizations: NewOrganizationStore(db, jwtMaker),
		Groups:        NewGroupStore(db),
		Time:          NewTimeStore(db),
		Instructors:   NewInstructorStore(db),
		Customers:     NewCustomerStore(db),
		Payments:      NewPaymentStore(db),
	}
}
