package cli

import "errors"

var (
	ErrMaxInstructionLength = errors.New("instructions must be less than 100 characters")
	ErrInvalidCommand       = errors.New("invalid command")
)
