package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
	"unsafe"
)

var (
	ErrInvalidDateFormat = errors.New("invalid date format. expected format: year-month-day")
)

type CustomDate struct {
	time.Time
}

func (cd CustomDate) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(unsafe.String(unsafe.SliceData(data), len(data)), `"`, "")
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return ErrInvalidDateFormat
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Time.Format(time.DateOnly))
}
