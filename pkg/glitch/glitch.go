package glitch

import "errors"

var (
	ErrItemAlreadyExists    = errors.New("item already exists in this collection")
	ErrUnknownDatabaseError = errors.New("unknown database error")
)
