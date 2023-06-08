package tests

import (
	"database/sql"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"testing"
)

func nullStrings(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func createRandomSuppliers(t *testing.T) db.Supplier {

	arg := db.CreateSupplierParams{
		Name:          utils.RandomName(),
		Address:       utils.RandomAddress(),
		ContactNumber: utils.RandomContact(),
	}

}
