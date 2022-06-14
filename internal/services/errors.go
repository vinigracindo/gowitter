package services

import "errors"

var (
	ErrUsernameMustBeUnique = errors.New("user with this username already exists")
)
