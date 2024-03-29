package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomProduct(t *testing.T, department db.Department, supplier db.Supplier) db.Product {
	arg := db.CreateProductsParams{
		Name:           utils.RandomName(),
		Description:    utils.RandomDesc(),
		Category:       utils.RandomAnyString(),
		DepartmentID:   department.DepartmentID,
		SupplierID:     supplier.SupplierID,
		Cost:           utils.RandomAnyInt(),
		SellingPrice:   utils.RandomAnyInt(),
		WholesalePrice: utils.RandomAnyInt(),
		MinMargin:      utils.RandomFloat(),
		Quantity:       utils.RandomAnyInt(),
	}
	product, err := testQueries.CreateProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Category, product.Category)
	require.Equal(t, arg.DepartmentID, product.DepartmentID)
	require.Equal(t, arg.SupplierID, product.SupplierID)
	require.Equal(t, arg.Cost, product.Cost)
	require.Equal(t, arg.SellingPrice, product.SellingPrice)
	require.Equal(t, arg.WholesalePrice, product.WholesalePrice)
	require.Equal(t, arg.MinMargin, product.MinMargin)
	require.Equal(t, arg.Quantity, product.Quantity)

	require.NotZero(t, product.ProductID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	department := createRandomDepartment(t)
	supplier := createRandomSuppliers(t)
	createRandomProduct(t, department, supplier)
}

func TestGetProduct(t *testing.T) {
	department := createRandomDepartment(t)
	supplier := createRandomSuppliers(t)

	product1 := createRandomProduct(t, department, supplier)

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

	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestUpdateProduct(t *testing.T) {
	department := createRandomDepartment(t)
	supplier := createRandomSuppliers(t)

	product1 := createRandomProduct(t, department, supplier)

	arg := db.UpdateProductsParams{
		ProductID:      product1.ProductID,
		Name:           utils.RandomName(),
		Description:    utils.RandomDesc(),
		Category:       utils.RandomAnyString(),
		Cost:           utils.RandomAnyInt(),
		SellingPrice:   utils.RandomAnyInt(),
		WholesalePrice: utils.RandomAnyInt(),
		MinMargin:      utils.RandomFloat(),
		Quantity:       utils.RandomAnyInt(),
	}

	product2, err := testQueries.UpdateProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ProductID, product2.ProductID)
	require.Equal(t, arg.Name, product2.Name)
	require.Equal(t, arg.Description, product2.Description)
	require.Equal(t, arg.Category, product2.Category)
	require.Equal(t, arg.Cost, product2.Cost)
	require.Equal(t, arg.SellingPrice, product2.SellingPrice)
	require.Equal(t, arg.WholesalePrice, product2.WholesalePrice)
	require.Equal(t, arg.MinMargin, product2.MinMargin)
	require.Equal(t, arg.Quantity, product2.Quantity)

	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestListProducts(t *testing.T) {
	var lastProduct db.Product

	for i := 0; i < 10; i++ {
		department := createRandomDepartment(t)
		supplier := createRandomSuppliers(t)

		product1 := createRandomProduct(t, department, supplier)

		lastProduct = product1
	}

	arg := db.ListProductsParams{
		ProductID: lastProduct.ProductID,
		Limit:     5,
		Offset:    0,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, products)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Equal(t, lastProduct.ProductID, product.ProductID)
	}
}

func TestDeleteProduct(t *testing.T) {
	department := createRandomDepartment(t)
	supplier := createRandomSuppliers(t)

	product1 := createRandomProduct(t, department, supplier)

	err := testQueries.DeleteProducts(context.Background(), product1.ProductID)
	require.NoError(t, err)

	product, err := testQueries.GetProducts(context.Background(), product1.ProductID)
	require.Error(t, err)

	require.EqualError(t, err, utils.ErrRecordNotFound.Error())
	require.Empty(t, product)
}
