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

	return product
}
