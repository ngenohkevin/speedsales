package tests

import (
	db "github.com/ngenohkevin/speedsales/db/sqlc"
	"github.com/ngenohkevin/speedsales/utils"
	"testing"
)

func createCodeTranslator(t *testing.T) db.CodeTranslator {
	arg := db.CreateCodeTranslatorParams{
		MasterCode: utils.RandomAnyString(),
		LinkCode:   utils.RandomAnyString(),
		PkgQty:     utils.RandomFloat(),
		Discount:   utils.RandomFloat(),
	}
}
