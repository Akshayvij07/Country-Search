package errors

import "errors"

var (
	ErrKeyNotFound    = errors.New("key not found")
	ErrParams         = errors.New("invalid parameters")
	ErrInvalidCountry = errors.New("invalid country")
)
