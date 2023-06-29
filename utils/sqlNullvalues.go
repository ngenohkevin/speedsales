package utils

import (
	"database/sql"
	"encoding/json"
	"github.com/tabbed/pqtype"
)

func NullStrings(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func NullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: value,
		Valid: true,
	}
}
func NullRawMessage(value json.RawMessage) pqtype.NullRawMessage {
	return pqtype.NullRawMessage{
		RawMessage: value,
		Valid:      true,
	}
}

func NullBool(value bool) sql.NullBool {
	return sql.NullBool{
		Bool:  value,
		Valid: true,
	}
}
