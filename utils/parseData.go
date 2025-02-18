package utils

import (
	"fmt"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	parsed_date, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Print(err)
		return time.Time{}, err
	}
	return parsed_date, nil
}
