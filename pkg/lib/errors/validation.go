package errors

import "fmt"

const (
	ValidationRequiredField  = "required"
	ValidationInvalidValue   = "invalid_value"
	ValidationValueNotUnique = "value_not_unique"
)

type ValidationErr struct {
	code  string
	field string
}

func NewValidation(field, code string) ValidationErr {
	return ValidationErr{
		field: field,
		code:  code,
	}
}

func (e ValidationErr) Code() string {
	return e.code
}

func (e ValidationErr) Field() string {
	return e.field
}

func (e ValidationErr) Error() string {
	return fmt.Sprintf("%s: %s", e.field, e.code)
}
