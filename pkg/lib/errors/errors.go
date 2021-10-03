package errors

import (
	"errors"
	"fmt"
)

// Error codes
// client may see these errors, so they should not contain sensitive data.
// nolint:gochecknoglobals // default errors
var (
	NotFound         = New("not found")
	AlreadyExists    = New("already exists")
	InternalError    = New("internal error")
	Unauthorized     = New("unauthorized")
	BadRequest       = New("bad request")
	PermissionDenied = New("permission denied")
	NoConnection     = New("no connection")
	Validation       = New("validation")
)

func New(code string) Err {
	return Err{
		code: code,
	}
}

// Err is a client error type.
type Err struct {
	code string
	err  error
}

func (e Err) Code() string {
	return e.code
}

func (e Err) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.code, e.err)
	}

	return e.code
}

func (e Err) Wrap(err error) Err {
	e.err = err

	return e
}

func (e Err) Unwrap() error {
	return e.err
}

func (e Err) Is(err error) bool {
	var codeErr Err
	if errors.As(err, &codeErr) {
		return codeErr.code == e.code
	}

	return false
}
