package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCodeTranslator(t *testing.T) db.CodeTranslator {
	arg := db.CreateCodeTranslatorParams{
		MasterCode: utils.RandomAnyString(),
		LinkCode:   utils.RandomAnyString(),
		PkgQty:     utils.RandomFloat(),
		Discount:   utils.RandomFloat(),
	}
	codeTranslator, err := testQueries.CreateCodeTranslator(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, codeTranslator)

	require.Equal(t, arg.MasterCode, codeTranslator.MasterCode)
	require.Equal(t, arg.LinkCode, codeTranslator.LinkCode)
	require.Equal(t, arg.PkgQty, codeTranslator.PkgQty)
	require.Equal(t, arg.Discount, codeTranslator.Discount)

	return codeTranslator
}

func TestCreateCodeTranslator(t *testing.T) {
	createRandomCodeTranslator(t)
}

func TestGetCreateCodeTranslator(t *testing.T) {
	codeTranslator1 := createRandomCodeTranslator(t)

	codeTranslator2, err := testQueries.GetCodeTranslator(context.Background(), codeTranslator1.MasterCode)
	require.NoError(t, err)
	require.NotEmpty(t, codeTranslator2)

	require.Equal(t, codeTranslator1.MasterCode, codeTranslator2.MasterCode)
	require.Equal(t, codeTranslator1.LinkCode, codeTranslator2.LinkCode)
	require.Equal(t, codeTranslator1.PkgQty, codeTranslator2.PkgQty)
	require.Equal(t, codeTranslator1.Discount, codeTranslator2.Discount)
}

func TestUpdateCodeTranslator(t *testing.T) {
	codeTranslator1 := createRandomCodeTranslator(t)

	arg := db.UpdateCodeTranslatorParams{
		MasterCode: codeTranslator1.MasterCode,
		PkgQty:     utils.RandomFloat(),
		Discount:   utils.RandomFloat(),
	}

	codeTranslator2, err := testQueries.UpdateCodeTranslator(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, codeTranslator2)

	require.Equal(t, codeTranslator1.MasterCode, codeTranslator2.MasterCode)
	require.Equal(t, arg.MasterCode, codeTranslator2.MasterCode)
	require.Equal(t, arg.MasterCode, codeTranslator2.MasterCode)
}

func TestListCodeTranslator(t *testing.T) {
	var lastCodeTranslator db.CodeTranslator

	for i := 0; i < 10; i++ {
		lastCodeTranslator = createRandomCodeTranslator(t)
	}
	arg := db.ListCodeTranslatorParams{
		MasterCode: lastCodeTranslator.MasterCode,
		Limit:      5,
		Offset:     0,
	}

	codeTranslator, err := testQueries.ListCodeTranslator(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, codeTranslator)

	for _, translator := range codeTranslator {
		require.NotEmpty(t, translator)
		require.Equal(t, lastCodeTranslator.MasterCode, translator.MasterCode)
	}
}
