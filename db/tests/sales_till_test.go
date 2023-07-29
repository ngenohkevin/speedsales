package tests

import (
	"context"
	"encoding/json"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomSalesTill(t *testing.T) db.SalesTill {
	arg := db.CreateSales_tillParams{
		TillNum:      utils.RandomAnyInt(),
		Teller:       utils.NullStrings(utils.RandomAnyString()),
		Supervisor:   utils.NullStrings(utils.RandomAnyString()),
		Branch:       utils.NullStrings(utils.RandomAnyString()),
		CloseTime:    utils.NullTimeStamp(time.Date(2023, time.July, 27, 11, 18, 10, 0, time.Local)), // Set the fixed CloseTime value
		CloseCash:    utils.NullFloat64(utils.RandomFloat()),
		CloseSummary: utils.RandomJSON(4),
	}
	salesTill, err := testQueries.CreateSales_till(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, salesTill)

	//Assert the CloseSummary field because it's JSON.
	//The tests will fail because of formatting and spacing at the CloseSummary field
	//The solution is to create two empty map representation which will hold the unmarshalled JSON objects
	//Unmarshall the RawMessage field of arg.CloseSummary and users.CloseSummary which allows conversion of JSON strings to map variables
	// require.Equal is used to compare the expectedCloseSummary with actualCloseSummary.
	//This assertion checks the unmarshalled JSON objects are equal regardless of their string representation
	//By comparing maps instead of raw JSON string, differences caused by formatting and or spacing is eliminated ensuring a valid comparison
	expectedCloseSummary := make(map[string]interface{})
	actualCloseSummary := make(map[string]interface{})
	err = json.Unmarshal(arg.CloseSummary, &expectedCloseSummary)
	require.NoError(t, err)
	err = json.Unmarshal(salesTill.CloseSummary, &actualCloseSummary)
	require.NoError(t, err)

	//require.Equal to compare the expected vs the actual values
	require.Equal(t, arg.TillNum, salesTill.TillNum)
	require.Equal(t, arg.Teller, salesTill.Teller)
	require.Equal(t, arg.Supervisor, salesTill.Supervisor)
	require.Equal(t, arg.Branch, salesTill.Branch)
	require.Equal(t, arg.CloseCash, salesTill.CloseCash)
	require.Equal(t, expectedCloseSummary, actualCloseSummary)

	return salesTill
}

func TestCreateSalesTill(t *testing.T) {
	createRandomSalesTill(t)
}

func TestGetSalesTill(t *testing.T) {
	salesTill1 := createRandomSalesTill(t)

	salesTill2, err := testQueries.GetSales_till(context.Background(), salesTill1.TillNum)
	require.NoError(t, err)
	require.NotEmpty(t, salesTill2)

	expectedCloseSummary := make(map[string]interface{})
	actualCloseSummary := make(map[string]interface{})
	err = json.Unmarshal(salesTill1.CloseSummary, &expectedCloseSummary)
	require.NoError(t, err)
	err = json.Unmarshal(salesTill2.CloseSummary, &actualCloseSummary)
	require.NoError(t, err)

	require.Equal(t, salesTill1.TillNum, salesTill2.TillNum)
	require.Equal(t, salesTill1.Teller, salesTill2.Teller)
	require.Equal(t, salesTill1.Supervisor, salesTill2.Supervisor)
	require.Equal(t, salesTill1.Branch, salesTill2.Branch)
	require.Equal(t, salesTill1.CloseCash, salesTill2.CloseCash)
	require.Equal(t, expectedCloseSummary, actualCloseSummary)

}
func TestUpdateSalesTill(t *testing.T) {

}
