// Package "glock" provides the core glock functionality, such as performing locks and unlocks.
package glock

import (
	"github.com/KyleBanks/glock/src/log"
	"sync"
	"time"
	"github.com/satori/go.uuid"
)

var (
	// noExpire is used when a lock is not created with an expire time, and
	// is used to indicate that the lock will never expire.
	noExpire = time.Time{}
)

type glock struct {
	mu *sync.Mutex
	lockRegister map[string]*lock

	logger *log.Logger
}

type lock struct {
	secret string
	expire time.Time
}

// New instantiates and returns a glock instance.
func New(logger *log.Logger) *glock {
	return &glock{
		mu: &sync.Mutex{},
		lockRegister: make(map[string]*lock),

		logger: logger,
	}
}

// lock attempts to register a lock on a key with the specified expire time.
func (g glock) lock(key string, expire time.Time) (string, *GlockError) {
	g.logger.Printf("lock(%v, %v)", key, expire)

	// Check if we're already locked
	if g.isLocked(key) {
		return "", ErrLockExists
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	// Looks good, create the new lock
	lock := &lock{
		secret: g.generateSecret(),
		expire: expire,
	}
	g.lockRegister[key] = lock

	return lock.secret, nil
}

// unlock immediately removes the lock on a key.
// This method is expected to be called after validation and after the glock mutex is locked.
func (g glock) unlock(key string) {
	delete(g.lockRegister, key)
}

// isLocked returns a boolean indicating if a key is locked, and ensures that if a lock exists,
// the expire time has not passed.
func (g glock) isLocked(key string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if lock, ok := g.lockRegister[key]; !ok {
		return false
	} else if lock.expire != noExpire && lock.expire.UnixNano() <= time.Now().UnixNano() {
		g.unlock(key)
		return false
	}
	return true
}

// secretMatches compares the secret provided with the secret stored alongside the lock of the key.
func (g glock) secretMatches(key, secret string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	lock := g.lockRegister[key];
	return lock.secret == secret
}

// Lock attempts to register a lock on the specified key and return a secret value that
// can later be used to unlock.
func (g glock) Lock(key string) (string, *GlockError) {
	g.logger.Printf("Lock(%v)", key)
	return g.lock(key, noExpire)
}

// LockWithDuration is the same as Lock, however it automatically expires the key after
// a specified duration.
func (g glock) LockWithDuration(key string, durationMs int) (string, *GlockError) {
	g.logger.Printf("LockWithDuration(%v, %v)", key, durationMs)

	// Calculate the expire time
	expire := time.Now().Add(time.Millisecond * time.Duration(durationMs))

	// Create the lock
	secret, err := g.lock(key, expire)
	if err != nil {
		return "", err
	}

	// Unlock the key after the specified duration
	go func(key, secret string, durationMs int) {
		time.Sleep(expire.Sub(time.Now()))
		g.Unlock(key, secret)
	}(key, secret, durationMs)

	return secret, nil
}

// Unlock attempts to unlock a locked key using the secret provided.
// If the secret doesn't match, or the key isn't locked, an error is returned.
func (g glock) Unlock(key, secret string) *GlockError {
	g.logger.Printf("Unlock(%v, %v)", key, secret)

	// If it's not locked, or the secret doesn't match, return
	if !g.isLocked(key) {
		return ErrLockNotExists
	} else if !g.secretMatches(key, secret) {
		return ErrSecretDoesNotMatch
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	// Looks good, unlock
	g.unlock(key)

	return nil
}

// Extend attempts to extend the lock on a key.
// If the key was locked without an expire time, the extension will be the current time + durationMs.
// If the key isn't locked or the secret doesn't match, an error is returned.
func (g glock) Extend(key, secret string, durationMs int) *GlockError {
	g.logger.Printf("Extend(%v, %v, %v)", key, secret, durationMs)

	// Validate the key and secret
	if !g.isLocked(key) {
		return ErrLockNotExists
	} else if !g.secretMatches(key, secret) {
		return ErrSecretDoesNotMatch
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	// Looks good, extend the lock
	lock := g.lockRegister[key]
	if lock.expire == noExpire {
		lock.expire = time.Now()
	}
	lock.expire = lock.expire.Add(time.Millisecond * time.Duration(durationMs))

	return nil
}

// generateSecret creates and returns a unique secret key.
func (g glock) generateSecret() string {
	return uuid.NewV4().String()
}