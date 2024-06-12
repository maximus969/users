package domain

import "errors"

var (
	ErrRequired      = errors.New("required value")
	ErrNotFound      = errors.New("not found")
	ErrNil           = errors.New("nil data")
	ErrNegative      = errors.New("negative value")
	ErrInvalidUserID = errors.New("invalid user ID")
)
