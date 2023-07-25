package tests

import (
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"testing"
	"time"
)

func createRandomSalesTill(t *testing.T) db.SalesTill {
	arg := db.CreateSales_tillParams{
		TillNum:      utils.RandomAnyInt(),
		Teller:       utils.NullStrings(utils.RandomAnyString()),
		Supervisor:   utils.NullStrings(utils.RandomAnyString()),
		Branch:       utils.NullStrings(utils.RandomAnyString()),
		CloseTime:    utils.NullTimeStamp(time.Now()),
		CloseCash:    utils.NullFloat64(utils.RandomFloat()),
		CloseSummary: utils.RandomJSON(4),
	}

}
