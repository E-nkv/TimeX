package api

import "fmt"

var (
	ErrDeadlineExceeded = fmt.Errorf("deadline exceeded")
	ErrInvalidArguments = fmt.Errorf("invalid args")
	ErrNotFound         = fmt.Errorf("entity not found")
)
