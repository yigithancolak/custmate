package postgresdb

import (
	"github.com/jmoiron/sqlx"
	"github.com/yigithancolak/custmate/util"
)

func NewDB(config *util.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", config.PostgresURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
