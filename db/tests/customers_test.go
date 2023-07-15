package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
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

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}
