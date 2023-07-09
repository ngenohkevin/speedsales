package tests

import (
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
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
}
