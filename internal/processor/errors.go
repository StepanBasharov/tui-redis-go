package processor

import "errors"

// Sentinel errors returned by command builders during input validation.
var (
	ErrUnknownCommand    = errors.New("unknown command")
	ErrNotEnoughArgs     = errors.New("not enough arguments")
	ErrInvalidSetOptions = errors.New("invalid SET option")
)
