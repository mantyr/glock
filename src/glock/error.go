package glock

import (
	"fmt"
)

var (
	ErrUnknown = &GlockError{code: 0, message: "An unexpected error occurred. Please check the logs."}

	// Returned when a key is not provided.
	ErrMissingKey = &GlockError{code: 1, message: "Missing required param: 'key'"}

	// Returned when a secret is not provided.
	ErrMissingSecret = &GlockError{code: 2, message: "Missing required param: 'secret'"}

	// Returned when an invalid duration is provided.
	ErrInvalidDuration = &GlockError{code: 3, message: "Invalid 'duration' provided."}

	// Returned when attempting to place a lock on a key that is already locked.
	ErrLockExists = &GlockError{code: 4, message: "Key already locked."}

	// Returned when attempting to perform an action on a key that is not locked.
	ErrLockNotExists = &GlockError{code: 5, message: "Key is not locked."}

	// Returned when attempting to use a secret that doesn't match the locked key's secret.
	ErrSecretDoesNotMatch = &GlockError{code: 6, message: "Secret does not match."}


)

type GlockError struct {
	code int `json:"code"`
	message string `json:"message"`
}

func (e GlockError) String() string {
	return fmt.Sprintf(`{code: %v, message: %v}`, e.code, e.message)
}