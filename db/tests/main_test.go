package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	testQueries = db.New(connPool)
	os.Exit(m.Run())

}
