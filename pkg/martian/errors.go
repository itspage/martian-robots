package martian

import "errors"

var (
	ErrMaxGridWidth  = errors.New("grid width must be less than 50")
	ErrMaxGridHeight = errors.New("grid height must be less than 50")
	ErrMinGridWidth  = errors.New("grid width must be greater than 0")
	ErrMinGridHeight = errors.New("grid height must be greater than 0")
)
