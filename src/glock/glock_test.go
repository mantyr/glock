package glock

import (
	"testing"
	"fmt"
	"time"
	"github.com/KyleBanks/glock/src/log"
)

func TestGlock_Lock(t *testing.T) {
	g := New(log.New(false))

	// Valid case
	validKey := fmt.Sprintf("testLock:%v", time.Now().Unix())

	if secret, err := g.Lock(validKey); err != nil {
		t.Fatal(err)
	} else if len(secret) == 0 {
		t.Fatalf("Invalid empty secret returned")
	}

	// Already locked
	if secret, err := g.Lock(validKey); err != ErrLockExists {
		t.Fatalf("Expected ErrLockExists, got: %v", err)
	} else if len(secret) != 0 {
		t.Fatalf("Invalid secret returned: %v", secret)
	}
}

func TestGlock_LockWithDuration(t *testing.T) {
	g := New(log.New(false))

	// Valid case
	validKey := fmt.Sprintf("testLockWithDuration:%v", time.Now().Unix())
	timeout := 200
	if secret, err := g.LockWithDuration(validKey, timeout); err != nil {
		t.Fatal(err)
	} else if len(secret) == 0 {
		t.Fatalf("Invalid empty secret returned")
	}

	// Already locked
	if secret, err := g.LockWithDuration(validKey, timeout); err != ErrLockExists {
		t.Fatalf("Expected ErrLockExists, got: %v", err)
	} else if len(secret) != 0 {
		t.Fatalf("Invalid secret returned: %v", secret)
	}

	// Wait for the lock to expire
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	// Try to lock again, expecting success
	if secret, err := g.LockWithDuration(validKey, timeout); err != nil {
		t.Fatal(err)
	} else if len(secret) == 0 {
		t.Fatalf("Invalid empty secret returned")
	}
}

func TestGlock_Unlock(t *testing.T) {
	g := New(log.New(false))

	// Not locked
	if err := g.Unlock("not locked", "secret"); err != ErrLockNotExists {
		t.Fatalf("Expected ErrLockNotExists, got: %v", err)
	}

	// Bad secret
	badSecretKey := fmt.Sprintf("testBadSecretKey:%v", time.Now().Unix())
	g.Lock(badSecretKey)
	if err := g.Unlock(badSecretKey, "bad secret"); err != ErrSecretDoesNotMatch {
		t.Fatalf("Expected ErrSecretDoesNotMatch, got: %v", err)
	}

	// Valid
	validKey := fmt.Sprintf("testUnlock:%v", time.Now().Unix())
	validSecret, _ := g.Lock(validKey)
	if err := g.Unlock(validKey, validSecret); err != nil {
		t.Fatal(err)
	}
}

func TestGlock_Unlock_LockWithDuration(t *testing.T) {
	g := New(log.New(false))

	// Valid
	validKey := fmt.Sprintf("testUnlockWithDuration:%v", time.Now().Unix())
	validSecret, _ := g.LockWithDuration(validKey, 20000)
	if err := g.Unlock(validKey, validSecret); err != nil {
		t.Fatal(err)
	}
}

func TestGlock_Extend(t *testing.T) {
	g := New(log.New(false))

	// Not locked
	if err := g.Extend("not locked", "secret", 1000); err != ErrLockNotExists {
		t.Fatalf("Expected ErrLockNotExists, got: %v", err)
	}

	// Bad secret
	badSecretKey := fmt.Sprintf("testBadSecretKeyExtend:%v", time.Now().Unix())
	g.Lock(badSecretKey)
	if err := g.Extend(badSecretKey, "bad secret", 1000); err != ErrSecretDoesNotMatch {
		t.Fatalf("Expected ErrSecretDoesNotMatch, got: %v", err)
	}

	// Valid
	validKey := fmt.Sprintf("testExtend:%v", time.Now().Unix())
	validSecret, _ := g.Lock(validKey)
	duration := 1000
	if err := g.Extend(validKey, validSecret, duration); err != nil {
		t.Fatal(err)
	}

	// Ensure the extension is removed after the current time * the durationProvided above
	time.Sleep(time.Duration(duration) * time.Millisecond)
	if _, err := g.Lock(validKey); err != nil {
		t.Fatal(err)
	}
}

func TestGlock_Extend_LockWithDuration(t *testing.T) {
	g := New(log.New(false))

	// Valid
	validKey := fmt.Sprintf("testExtendWithDuration:%v", time.Now().Unix())
	duration := 1000
	validSecret, _ := g.LockWithDuration(validKey, duration)

	// Extend within the duration
	time.Sleep(time.Duration(duration / 2) * time.Millisecond)
	if err := g.Extend(validKey, validSecret, duration); err != nil {
		t.Fatal(err)
	}

	// Now wait passed the duration and ensure it can't be extended (because its no longer locked)
	time.Sleep(time.Duration(duration * 2) * time.Millisecond)
	if err := g.Extend(validKey, validSecret, duration); err != ErrLockNotExists {
		t.Fatal(err)
	}
}