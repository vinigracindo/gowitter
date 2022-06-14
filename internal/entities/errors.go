package entities

import "errors"

var (
	ErrAlreadyFollowing     = errors.New("already following")
	ErrCannotFollowYourself = errors.New("cannot follow yourself")
)
