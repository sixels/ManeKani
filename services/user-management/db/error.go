package db

import (
	"errors"
	"fmt"
)

var (
	ErrUnknown = errors.New("unknown database error")
)

func unknownError(err error) error {
	return fmt.Errorf("%w: %w", ErrUnknown, err)
}
