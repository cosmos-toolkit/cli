package entity

import "errors"

var (
	// ErrNotFound is returned when a resource does not exist.
	ErrNotFound = errors.New("entity: not found")
	// ErrInvalidInput is returned when the input is invalid.
	ErrInvalidInput = errors.New("entity: invalid input")
	// ErrConflict is returned when there is a conflict (e.g. duplicate).
	ErrConflict = errors.New("entity: conflict")
)
