package models

import "errors"

var (
	ErrNoRecords = errors.New("models: no matching records found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)
