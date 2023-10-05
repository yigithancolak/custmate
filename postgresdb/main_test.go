package postgresdb

import (
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/yigithancolak/custmate/token"
	"github.com/yigithancolak/custmate/util"
)

type StoreTestSuite struct {
	suite.Suite
	store *Store
}

func (s *StoreTestSuite) SetupSuite() {
	config, err := util.LoadConfig("..", "test", "env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := sqlx.Connect("pgx", config.ConnectionURL)
	if err != nil {
		s.T().Log(err.Error())
	}
	_, err = db.Exec("CREATE DATABASE test_custmate")
	if err != nil {
		log.Fatal("cannot create test db:", err)
	}
	db.Close()

	testDB, err := NewDB(config)
	if err != nil {
		log.Fatal("cannot create test db instance:", err)
	}

	jwtMaker, err := token.NewJWTMaker(config)
	if err != nil {
		log.Fatal("cannot creating jwt maker:", err)
	}

	m, err := migrate.New(
		"file://./migration",
		config.MigrationURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	s.store = NewStore(testDB, jwtMaker)
}
func (s *StoreTestSuite) TearDownSuite() {
	config, err := util.LoadConfig("..", "test", "env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	s.store.DB.Close() // Close the testDB connection

	// Connect to the default database (or any other database that's not the one you're trying to drop)
	db, err := sqlx.Connect("pgx", config.ConnectionURL)
	if err != nil {
		s.T().Log(err.Error())
		return
	}
	defer db.Close()

	// Forcefully terminate all connections to the test_custmate database
	_, err = db.Exec("SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'test_custmate';")
	if err != nil {
		s.T().Logf("Error terminating connections to test database: %v", err)
	}

	// Drop the test_custmate database
	_, err = db.Exec("DROP DATABASE IF EXISTS test_custmate")
	if err != nil {
		s.T().Logf("Error dropping test database: %v", err)
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
