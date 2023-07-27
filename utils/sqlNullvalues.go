package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func NullStrings(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func NullInt64(value int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: value,
		Valid: true,
	}
}

//func NullRawMessage(value json.RawMessage) pqtype.NullRawMessage {
//	return pqtype.NullRawMessage{
//		RawMessage: value,
//		Valid:      true,
//	}
//}

func NullBool(value bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  value,
		Valid: true,
	}
}
func NullTimeStamp(value time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:             value,
		InfinityModifier: 0,
		Valid:            true,
	}
}

func NullFloat64(value float64) pgtype.Float8 {
	return pgtype.Float8{
		Float64: value,
		Valid:   true,
	}
}
