package types

import (
	"fmt"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalText(text []byte) error {
	parsedTime, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	*d = Date(parsedTime)
	return nil
}
