package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// Helper function to round the time to the nearest microsecond
func roundTime(t time.Time, precision time.Duration) time.Time {
	rounded := t.Round(precision)
	// Ensure the rounded time is rounded to the nearest microsecond
	return time.Unix(rounded.Unix(), int64(rounded.Nanosecond()/1000)*1000)
}

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
	salesTill, err := testQueries.CreateSales_till(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, salesTill)

	currentTimeStamp := roundTime(time.Now(), time.Millisecond)
	expectedCloseTime := roundTime(currentTimeStamp, time.Millisecond)
	actualCloseTime := roundTime(salesTill.CloseTime.Time, time.Millisecond)

	require.Equal(t, arg.TillNum, salesTill.TillNum)
	require.Equal(t, arg.Teller, salesTill.Teller)
	require.Equal(t, arg.Supervisor, salesTill.Supervisor)
	require.Equal(t, arg.Branch, salesTill.Branch)
	require.Equal(t, expectedCloseTime, actualCloseTime)
	require.Equal(t, arg.CloseCash, salesTill.CloseCash)
	require.Equal(t, arg.CloseSummary, salesTill.CloseSummary)

	return salesTill
}

func TestCreateSalesTill(t *testing.T) {
	createRandomSalesTill(t)
}
