package utils

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

func ConvertDate(timestamp *time.Time) datatypes.Date {
	layout := "2006-01-02"
	date := fmt.Sprintf("%d-%d-%d", timestamp.Year(), timestamp.Month(), timestamp.Day())
	ddate, _ := time.Parse(layout, date)
	return datatypes.Date(ddate)
}
