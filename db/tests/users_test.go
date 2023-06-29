package tests

import (
	"context"
	"encoding/json"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUsers(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username:    utils.NullStrings(utils.RandomName()),
		Branch:      utils.NullStrings(utils.RandomAnyString()),
		StkLocation: utils.NullStrings(utils.RandomAnyString()),
		Reset:       utils.NullStrings(utils.RandomAnyString()),
		TillNum:     utils.NullInt64(int64(utils.RandomAnyInt())),
		Rights:      utils.NullRawMessage(utils.RandomJSON(2)),
		IsActive:    utils.NullBool(utils.RandomBool()),
	}
	users, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	//Assert the Rights field because it's JSON.
	//The tests will fail because of formatting and spacing at the Rights field
	//The solution is to create two empty map representation which will hold the unmarshalled JSON objects
	//Unmarshall the RawMessage field of arg.Rights and users.Rights which allows conversion of JSON strings to map variables
	// require.Equal is used to compare the expectedRights with actualRights.
	//This assertion checks the unmarshalled JSON objects are equal regardless of their string representation
	//By comparing maps instead of raw JSON string, differences caused by formatting and or spacing is eliminated ensuring a valid comparison
	expectedRights := make(map[string]interface{})
	actualRights := make(map[string]interface{})
	err = json.Unmarshal(arg.Rights.RawMessage, &expectedRights)
	require.NoError(t, err)
	err = json.Unmarshal(users.Rights.RawMessage, &actualRights)
	require.NoError(t, err)

	require.Equal(t, arg.Username, users.Username)
	require.Equal(t, arg.Branch, users.Branch)
	require.Equal(t, arg.StkLocation, users.StkLocation)
	require.Equal(t, arg.Reset, users.Reset)
	require.Equal(t, expectedRights, actualRights)
	require.Equal(t, arg.IsActive, users.IsActive)

	require.NotZero(t, users.UserID)

	return users
}

func TestCreateUsers(t *testing.T) {
	createRandomUsers(t)
}
