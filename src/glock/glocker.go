package glock

// The Glocker interface defines the functionality of a 'glock'.
// The interface allows for multiple implementations, such as a mock 'glock'.
type Glocker interface {
	// Lock attempts to register a lock on the specified key with a secret value that
	// can later be used to unlock.
	Lock(key, secret string) *GlockError

	// LockWithDuration is the same as Lock, however it automatically expires the key after
	// a specified duration.
	LockWithDuration(key, secret string, durationMs int) *GlockError

	// Unlock attempts to unlock a locked key using the secret provided.
	// If the secret doesn't match, or the key isn't locked, an error is returned.
	Unlock(key, secret string) *GlockError
}
