// Package "glock" provides the core glock functionality, such as performing locks and unlocks.
package glock

import "github.com/KyleBanks/glock/src/log"

type glock struct {
	LockRegister map[string]string

	Logger *log.Logger
}

type Glocker interface {
	// Lock attempts to register a lock on the specified key with a secret value that
	// can later be used to unlock.
	Lock(key, secret string) error

	// LockWithDuration is the same as Lock, however it automatically expires the key after
	// a specified duration.
	LockWithDuration(key, secret string, durationMs int) error
}

// New instantiates and returns a glock instance.
func New(logger *log.Logger) *glock {
	return &glock{
		Logger: logger,
	}
}

func (g glock) Lock(key, secret string) error {
	g.Logger.Printf("Lock(%v, %v)", key, secret)
	return nil
}

func (g glock) LockWithDuration(key, secret string, durationMs int) error {
	return nil
}
