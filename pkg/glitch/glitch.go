package glitch

import "errors"

var (
	ErrRecordNotFound       = errors.New("database record not found")
	ErrItemAlreadyExists    = errors.New("item already exists in this collection")
	ErrUnknownDatabaseError = errors.New("unknown database error")
)
