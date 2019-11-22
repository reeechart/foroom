package errors

import "errors"

var (
	ErrInvalidArgs = errors.New("Invalid user or room name")
)
