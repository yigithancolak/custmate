package postgresdb

import (
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/yigithancolak/custmate/token"
	"github.com/yigithancolak/custmate/util"
)

var testStore *Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("..", "development", "env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := NewDB(config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	jwtMaker, err := token.NewJWTMaker(config)
	if err != nil {
		log.Fatal("cannot creating jwt maker:", err)
	}

	testStore = NewStore(db, jwtMaker)
	os.Exit(m.Run())

}
