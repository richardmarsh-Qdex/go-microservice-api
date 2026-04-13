package apperrors

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrConflict      = errors.New("conflict")
	ErrInsufficient  = errors.New("insufficient stock")
)
