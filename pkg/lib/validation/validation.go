package validation

import (
	"fmt"

	"github.com/lab1/errors-logs/pkg/lib/errors"
)

type ColumnString struct {
	// Name - column name
	Name string

	// Optional - if true allow empty value
	Optional bool

	// MaxLen - if 0 don't check value length
	MaxLen int
}

type ColumnStringSlice struct {
	// Name - column name
	Name string

	// Optional - if true allow empty value
	Optional bool

	// MaxLen - if 0 don't check value length
	MaxLen int

	// ValueMaxLen - if 0 don't check value length
	ValueMaxLen int
}

type ColumnInt struct {
	// Name - column name
	Name string

	// Optional - if true allow empty value
	Optional bool

	// MaxValue - if 0 don't check value length
	MaxValue int
}

func CheckString(c ColumnString, v string) error {
	if !c.Optional && v == "" {
		return errors.NewValidation(c.Name, errors.ValidationRequiredField)
	}

	if c.MaxLen > 0 && len(v) > c.MaxLen {
		return errors.NewValidation(c.Name, errors.ValidationInvalidValue)
	}

	return nil
}

func CheckStringSlice(c ColumnStringSlice, slice []string) error {
	if !c.Optional && len(slice) == 0 {
		return errors.NewValidation(c.Name, errors.ValidationRequiredField)
	}

	if c.MaxLen > 0 && len(slice) > c.MaxLen {
		return errors.NewValidation(c.Name, errors.ValidationInvalidValue)
	}

	return nil
}

func CheckInt(c ColumnInt, v int) error {
	if !c.Optional && v == 0 {
		return errors.NewValidation(c.Name, errors.ValidationRequiredField)
	}

	if c.MaxValue > 0 && v > c.MaxValue {
		return errors.NewValidation(c.Name, errors.ValidationInvalidValue)
	}

	return nil
}

func ValidateString(errb *errors.Bundle, c ColumnString, v string) {
	if err := CheckString(c, v); err != nil {
		errb.Add(err)
	}
}

func ValidateStringSlice(errb *errors.Bundle, c ColumnStringSlice, slice []string) {
	if err := CheckStringSlice(c, slice); err != nil {
		errb.Add(err)
	}

	field := ""
	for k, v := range slice {
		field = fmt.Sprintf("%s%d", c.Name, k)
		ValidateString(errb, ColumnString{Name: field, MaxLen: c.ValueMaxLen}, v)
	}
}

func ValidateInt(errb *errors.Bundle, c ColumnInt, v int) {
	if err := CheckInt(c, v); err != nil {
		errb.Add(err)
	}
}
