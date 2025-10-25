package utils

import "fmt"

var (
	ErrBadRequest = fmt.Errorf("bad request")
	ErrNotFound   = fmt.Errorf("not found")
	ErrForbidden  = fmt.Errorf("forbidden")
	ErrInternal   = fmt.Errorf("internal server error")
)
