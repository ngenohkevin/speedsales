package tests

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())

}
