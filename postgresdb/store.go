package postgresdb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/token"
)

type Store struct {
	DB       *sqlx.DB
	JWTMaker *token.JWTMaker
}

type queryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func NewStore(db *sqlx.DB, jwtMaker *token.JWTMaker) *Store {
	return &Store{
		DB:       db,
		JWTMaker: jwtMaker,
	}
}
