package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

func TestGetProduct(t *testing.T) {
	supplier1 := createRandomSuppliers(t)

	supplier2, err := testQueries.GetSupplier(context.Background(), supplier1.SupplierID)
	require.NoError(t, err)
	require.NotEmpty(t, supplier2)

	require.Equal(t, supplier1.Name, supplier2.Name)
	require.Equal(t, supplier1.Address, supplier2.Address)
	require.Equal(t, supplier1.ContactNumber, supplier2.ContactNumber)
	require.Equal(t, supplier1.Email, supplier2.Email)

	require.WithinDuration(t, supplier1.CreatedAt, supplier2.CreatedAt, time.Second)
}

func TestListSupplier(t *testing.T) {
	//var lastSupplier db.Supplier

	for i := 0; i < 10; i++ {
		createRandomSuppliers(t)
	}
	arg := db.ListSupplierParams{
		Limit:  5,
		Offset: 5,
	}

	supplier, err := testQueries.ListSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, supplier)
	require.Len(t, supplier, 5)

	for _, suppliers := range supplier {
		require.NotEmpty(t, suppliers)
		//require.Equal(t, lastSupplier.SupplierID, suppliers.SupplierID)
	}
}

func TestUpdateSupplier(t *testing.T) {
	supplier1 := createRandomSuppliers(t)

	arg := db.UpdateSupplierParams{
		SupplierID:    supplier1.SupplierID,
		Name:          utils.RandomName(),
		Address:       utils.RandomAddress(),
		ContactNumber: utils.RandomContact(),
		Email:         utils.RandomEmail(),
	}

}
