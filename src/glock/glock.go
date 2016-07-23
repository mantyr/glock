// Package "glock" provides the core glock functionality, such as performing locks and unlocks.
package glock

import (
	"github.com/KyleBanks/glock/src/log"
	"sync"
	"time"
	"github.com/satori/go.uuid"
)

type glock struct {
	mu *sync.Mutex
	lockRegister map[string]string

	logger *log.Logger
}

// New instantiates and returns a glock instance.
func New(logger *log.Logger) *glock {
	return &glock{
		mu: &sync.Mutex{},
		lockRegister: make(map[string]string),

		logger: logger,
	}
}

// Lock attempts to register a lock on the specified key and return a secret value that
// can later be used to unlock.
func (g glock) Lock(key string) (string, *GlockError) {
	g.logger.Printf("Lock(%v)", key)

	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.lockRegister[key]; !ok {
		secret := g.generateSecret()
		g.lockRegister[key] = secret
		return secret, nil
	}

	return "", ErrLockExists
}
// LockWithDuration is the same as Lock, however it automatically expires the key after
// a specified duration.
func (g glock) LockWithDuration(key string, durationMs int) (string, *GlockError) {
	g.logger.Printf("LockWithDuration(%v, %v)", key, durationMs)

	secret, err := g.Lock(key)
	if err != nil {
		return "", err
	}

	go func(key, secret string, durationMs int) {
		time.Sleep(time.Duration(durationMs) * time.Millisecond)
		g.Unlock(key, secret)
	}(key, secret, durationMs)

	return secret, nil
}

// Unlock attempts to unlock a locked key using the secret provided.
// If the secret doesn't match, or the key isn't locked, an error is returned.
func (g glock) Unlock(key, secret string) *GlockError {
	g.logger.Printf("Unlock(%v, %v)", key, secret)

	g.mu.Lock()
	defer g.mu.Unlock()

	if val, ok := g.lockRegister[key]; !ok {
		return ErrLockNotExists
	} else if val != secret {
		return ErrSecretDoesNotMatch
	}

	delete(g.lockRegister, key)
	return nil
}

// generateSecret creates and returns a unique secret key.
func (g glock) generateSecret() string {
	return uuid.NewV4().String()
}