package utils

import "fmt"

var (
	ErrBadRequest      = fmt.Errorf("bad request")
	ErrNotFound        = fmt.Errorf("not found")
	ErrForbidden       = fmt.Errorf("forbidden")
	ErrInternal        = fmt.Errorf("internal server error")
	ErrInvalidUsername = fmt.Errorf("invalid username")
	ErrInvalidID       = fmt.Errorf("invalid ID")
)
var (
	FixInvalidUsername = "Please provide a valid username."
	FixInvalidID       = "Please provide a valid ID."
)
