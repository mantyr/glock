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
	if err := g.Lock(validKey, "secret"); err != nil {
		t.Fatal(err)
	}

	// Already locked
	if err := g.Lock(validKey, "secret"); err != ErrLockExists {
		t.Fatalf("Expected ErrLockExists, got: %v", err)
	}
}

func TestGlock_LockWithDuration(t *testing.T) {
	g := New(log.New(false))

	// Valid case
	validKey := fmt.Sprintf("testLockWithDuration:%v", time.Now().Unix())
	timeout := 1000 * 2
	if err := g.LockWithDuration(validKey, "secret", timeout); err != nil {
		t.Fatal(err)
	}

	// Already locked
	if err := g.LockWithDuration(validKey, "secret", timeout); err != ErrLockExists {
		t.Fatalf("Expected ErrLockExists, got: %v", err)
	}

	// Wait for the lock to expire
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	// Try to lock again, expecting success
	if err := g.LockWithDuration(validKey, "secret", timeout); err != nil {
		t.Fatal(err)
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
	g.Lock(badSecretKey, "secret")
	if err := g.Unlock(badSecretKey, "bad secret"); err != ErrSecretDoesNotMatch {
		t.Fatalf("Expected ErrSecretDoesNotMatch, got: %v", err)
	}

	// Valid
	validKey := fmt.Sprintf("testUnlock:%v", time.Now().Unix())
	validSecret := "a secret"
	g.Lock(validKey, validSecret)
	if err := g.Unlock(validKey, validSecret); err != nil {
		t.Fatal(err)
	}
}