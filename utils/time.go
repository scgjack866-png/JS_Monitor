package utils

import (
	"fmt"
	"time"
)

type Time time.Time

// returns time.Now() no matter what!
func (t Time) UnmarshalJSON() ([]byte, error) {
	// you can now parse b as thoroughly as you want
	dateTime := fmt.Sprintf("%q", time.Time(t).Format("2006-01-02"))
	return []byte(dateTime), nil
}
