package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomProduct(t *testing.T) db.Product {
	arg := db.CreateProductsParams{
		Name:           utils.RandomName(),
		Description:    utils.RandomDesc(),
		Category:       utils.RandomAnyString(),
		DepartmentID:   int32(utils.RandomAnyInt()),
		SupplierID:     utils.RandomAnyInt(),
		Cost:           utils.RandomAnyInt(),
		SellingPrice:   utils.RandomAnyInt(),
		WholesalePrice: utils.RandomAnyInt(),
		MinMargin:      utils.RandomFloat(),
		Quantity:       utils.RandomAnyInt(),
	}
	product, err := testQueries.CreateProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.NotZero(t, product.ProductID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomUsers(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	product2, err := testQueries.GetProducts(context.Background(), product1.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ProductID, product2.ProductID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.Category, product2.Category)
	require.Equal(t, product1.DepartmentID, product2.DepartmentID)
	require.Equal(t, product1.SupplierID, product2.SupplierID)
	require.Equal(t, product1.Cost, product2.Cost)
	require.Equal(t, product1.SellingPrice, product2.SellingPrice)
	require.Equal(t, product1.WholesalePrice, product2.WholesalePrice)
	require.Equal(t, product1.MinMargin, product2.MinMargin)
	require.Equal(t, product1.Quantity, product2.Quantity)
}
