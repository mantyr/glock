package glock

import (
	"fmt"
)

var (
	ErrUnknown = &GlockError{Code: 0, Message: "An unexpected error occurred. Please check the logs."}

	// Returned when a key is not provided.
	ErrMissingKey = &GlockError{Code: 1, Message: "Missing required param: 'key'"}

	// Returned when a secret is not provided.
	ErrMissingSecret = &GlockError{Code: 2, Message: "Missing required param: 'secret'"}

	// Returned when an invalid duration is provided.
	ErrInvalidDuration = &GlockError{Code: 3, Message: "Invalid 'duration' provided."}

	// Returned when attempting to place a lock on a key that is already locked.
	ErrLockExists = &GlockError{Code: 4, Message: "Key already locked."}

	// Returned when attempting to perform an action on a key that is not locked.
	ErrLockNotExists = &GlockError{Code: 5, Message: "Key is not locked."}

	// Returned when attempting to use a secret that doesn't match the locked key's secret.
	ErrSecretDoesNotMatch = &GlockError{Code: 6, Message: "Secret does not match."}
)

type GlockError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func (e GlockError) String() string {
	return fmt.Sprintf(`{code: %v, message: %v}`, e.Code, e.Message)
}