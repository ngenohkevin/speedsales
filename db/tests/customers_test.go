package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomCustomer(t *testing.T) db.Customer {
	arg := db.CreateCustomerParams{
		Name:          utils.RandomName(),
		Address:       utils.NullStrings(utils.RandomAddress()),
		ContactNumber: utils.RandomContact(),
		Email:         utils.NullStrings(utils.RandomEmail()),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.Name, customer.Name)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.ContactNumber, customer.ContactNumber)
	require.Equal(t, arg.Email, customer.Email)

	require.NotZero(t, customer.CustomerID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.CustomerID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.CustomerID, customer2.CustomerID)
	require.Equal(t, customer1.Name, customer2.Name)
	require.Equal(t, customer1.Address, customer2.Address)
	require.Equal(t, customer1.ContactNumber, customer2.ContactNumber)
	require.Equal(t, customer1.Email, customer2.Email)

	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := db.UpdateCustomerParams{
		CustomerID:    customer1.CustomerID,
		Name:          utils.RandomName(),
		Address:       utils.NullStrings(utils.RandomAddress()),
		ContactNumber: utils.RandomContact(),
		Email:         utils.NullStrings(utils.RandomEmail()),
	}

	customer2, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.CustomerID, customer2.CustomerID)
	require.Equal(t, arg.Name, customer2.Name)
	require.Equal(t, arg.Address, customer2.Address)
	require.Equal(t, arg.ContactNumber, customer2.ContactNumber)
	require.Equal(t, arg.Email, customer2.Email)

	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}

func TestListCustomers(t *testing.T) {
	var lastCustomer db.Customer

	for i := 0; i < 10; i++ {
		lastCustomer = createRandomCustomer(t)
	}

	arg := db.ListCustomersParams{
		CustomerID: lastCustomer.CustomerID,
		Limit:      5,
		Offset:     0,
	}
	customers, err := testQueries.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)

	for _, customer := range customers {
		require.NotEmpty(t, customer)
		require.Equal(t, lastCustomer.CustomerID, customer.CustomerID)
	}
}
