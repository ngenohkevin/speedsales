package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

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

func TestGetSupplier(t *testing.T) {
	supplier1 := createRandomSuppliers(t)

	supplier2, err := testQueries.GetSupplier(context.Background(), supplier1.SupplierID)
	require.NoError(t, err)
	require.NotEmpty(t, supplier2)

	require.Equal(t, supplier1.SupplierID, supplier2.SupplierID)
	require.Equal(t, supplier1.Name, supplier2.Name)
	require.Equal(t, supplier1.Address, supplier2.Address)
	require.Equal(t, supplier1.ContactNumber, supplier2.ContactNumber)
	require.Equal(t, supplier1.Email, supplier2.Email)

	require.WithinDuration(t, supplier1.CreatedAt, supplier2.CreatedAt, time.Second)
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
	supplier2, err := testQueries.UpdateSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, supplier2)

	require.Equal(t, supplier1.SupplierID, supplier2.SupplierID)
	require.Equal(t, arg.Name, supplier2.Name)
	require.Equal(t, arg.Address, supplier2.Address)
	require.Equal(t, arg.ContactNumber, supplier2.ContactNumber)
	require.Equal(t, arg.Email, supplier2.Email)
}

func TestListSuppliers(t *testing.T) {
	var lastSupplier db.Supplier

	for i := 0; i < 10; i++ {
		lastSupplier = createRandomSuppliers(t)
	}
	arg := db.ListSuppliersParams{
		SupplierID: lastSupplier.SupplierID,
		Limit:      5,
		Offset:     0,
	}
	suppliers, err := testQueries.ListSuppliers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, suppliers)

	for _, supplier := range suppliers {
		require.NotEmpty(t, supplier)
		require.Equal(t, lastSupplier.SupplierID, supplier.SupplierID)
	}
}
func TestDeleteSupplier(t *testing.T) {
	supplier1 := createRandomSuppliers(t)

	err := testQueries.DeleteSupplier(context.Background(), supplier1.SupplierID)
	require.NoError(t, err)

	supplier2, err := testQueries.GetSupplier(context.Background(), supplier1.SupplierID)
	require.Error(t, err)

	require.EqualError(t, err, utils.ErrRecordNotFound.Error())

	require.Empty(t, supplier2)
}
