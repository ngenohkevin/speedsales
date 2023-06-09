package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

//func nullStrings(value string) sql.NullString {
//	return sql.NullString{
//		String: value,
//		Valid:  true,
//	}
//}

func createRandomSuppliers(t *testing.T) db.Supplier {

	arg := db.CreateSupplierParams{
		Name:          utils.RandomName(),
		Address:       utils.RandomAddress(),
		ContactNumber: utils.RandomContact(),
		Email:         utils.RandomEmail(),
	}
	supplier, err := testQueries.CreateSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, supplier)

	require.Equal(t, arg.Name, supplier.Name)
	require.Equal(t, arg.Address, supplier.Address)
	require.Equal(t, arg.ContactNumber, supplier.ContactNumber)
	require.Equal(t, arg.Email, supplier.Email)

	require.NotZero(t, supplier.SupplierID)
	require.NotZero(t, supplier.CreatedAt)

	return supplier

}

func TestCreateSupplier(t *testing.T) {
	createRandomSuppliers(t)
}
