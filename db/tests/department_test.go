package tests

import (
	"context"
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomDepartment(t *testing.T) db.Department {

	arg := db.CreateDepartmentParams{
		Category:    utils.NullStrings(utils.RandomAnyString()),
		SubCategory: utils.NullStrings(utils.RandomAnyString()),
		Description: utils.NullStrings(utils.RandomDesc()),
	}
	department, err := testQueries.CreateDepartment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, department)

	require.Equal(t, arg.Category, department.Category)
	require.Equal(t, arg.SubCategory, department.SubCategory)
	require.Equal(t, arg.Description, department.Description)

	require.NotZero(t, department.DepartmentID)

	return department
}

func TestCreateDepartment(t *testing.T) {
	createRandomDepartment(t)
}

func TestGetDepartment(t *testing.T) {
	department1 := createRandomDepartment(t)

	department2, err := testQueries.GetDepartment(context.Background(), department1.DepartmentID)
	require.NoError(t, err)
	require.NotEmpty(t, department2)

	require.Equal(t, department1.DepartmentID, department2.DepartmentID)
	require.Equal(t, department1.Category, department2.Category)
	require.Equal(t, department1.SubCategory, department2.SubCategory)
	require.Equal(t, department1.Description, department2.Description)

}
