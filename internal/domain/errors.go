package domain

import "errors"

var (
	ErrNotFound   = errors.New("entity not found")
	ErrEmptyField = errors.New("field is empty")
)

type EmptyFieldError struct {
	field string
}

func (e *EmptyFieldError) Error() string {
	return e.field + " cannot be empty"
}

func (e *EmptyFieldError) Field() string { return e.field }

func (e *EmptyFieldError) Is(err error) bool {
	return errors.Is(err, ErrEmptyField)
}
