package api

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/KyleBanks/glock/src/glock"
	"fmt"
	"time"
)

func TestGlockApi_HandleLock_MissingKey(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	g.HandleLock(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrMissingKey.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleLock_Positive(t *testing.T) {
	g := NewForTest()

	key := fmt.Sprintf("testkey:%v", time.Now().Unix())

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key="+key, nil)
	g.HandleLock(w, r)

	res := ValidateApiResponse(t, w)
	if secret := res.Extras["secret"]; len(secret) == 0 {
		t.Fatalf("Missing Secret in response: %v", res)
	}
}

func TestGlockApi_HandleLock_AlreadyLocked(t *testing.T) {
	g := NewForTest()

	key := fmt.Sprintf("testkey:%v", time.Now().Unix())

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key="+key, nil)

	// Lock the key
	g.HandleLock(w, r)
	ValidateApiResponse(t, w)

	// Attempt to lock again
	w = httptest.NewRecorder()
	g.HandleLock(w, r)
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrLockExists.Code {
		t.Fatalf("Expected ErrLockExists, got: %v", res.Error)
	}
}