package repository

import "errors"

var (
	ErrNotImplement = errors.New("not implemented")
)

type ErrNotImplemented struct {
	field string
}

func (e *ErrNotImplemented) Error() string {
	return e.field + " cannot be empty"
}

func (e *ErrNotImplemented) Field() string { return e.field }

func (e *ErrNotImplemented) Is(err error) bool {
	return errors.Is(err, ErrNotImplement)
}
