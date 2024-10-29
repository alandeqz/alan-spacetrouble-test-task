package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Gender int

const (
	Unknown Gender = 0
	Male    Gender = 1
	Female  Gender = 2
	Other   Gender = 3
)

const (
	UnknownGender = "Unknown"
	MaleGender    = "Male"
	FemaleGender  = "Female"
	OtherGender   = "Other"
)

func (g Gender) String() string {
	return [...]string{
		UnknownGender,
		MaleGender,
		FemaleGender,
		OtherGender,
	}[g]
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case MaleGender:
		*g = Male
	case FemaleGender:
		*g = Female
	case OtherGender:
		*g = Other
	default:
		*g = Unknown
	}

	return nil
}

func (g Gender) Value() (driver.Value, error) {
	return int(g), nil
}

func (g *Gender) Scan(value interface{}) error {
	if value == nil {
		*g = Unknown

		return nil
	}

	switch v := value.(type) {
	case int:
		*g = Gender(v)
	case int32:
		*g = Gender(v)
	case int64:
		*g = Gender(v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}
