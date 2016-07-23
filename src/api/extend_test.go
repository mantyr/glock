package api

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/KyleBanks/glock/src/glock"
	"fmt"
	"time"
)

func TestGlockApi_HandleExtend_MissingKey(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrMissingKey.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_MissingSecret(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=test", nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrMissingSecret.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_MissingDuration(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=test&secret=test", nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrMissingDuration.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_InvalidDuration(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=test&secret=test&duration=test", nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrInvalidDuration.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_NotLocked(t *testing.T) {
	g := NewForTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=notLocked&secret=notNeeded&duration=10", nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrLockNotExists.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_BadSecret(t *testing.T) {
	g := NewForTest()

	// First, lock a key
	key := fmt.Sprintf("testbadsecret:%v", time.Now().Unix())
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=" + key, nil)
	g.HandleLock(w, r)
	ValidateApiResponse(t, w)

	// Attempt to extend using a bad secret
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", fmt.Sprintf("?key=%v&secret=badsecret&duration=10", key), nil)
	g.HandleExtend(w, r)

	// Validate an error was returned
	res := ValidateApiResponseError(t, w)
	if res.Error.Code != glock.ErrSecretDoesNotMatch.Code {
		t.Fatalf("Unexpected error returned: %v", res.Error)
	}
}

func TestGlockApi_HandleExtend_Positive(t *testing.T) {
	g := NewForTest()

	// First, lock a key
	key := fmt.Sprintf("testbadsecret:%v", time.Now().Unix())
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?key=" + key, nil)
	g.HandleLock(w, r)
	lockRes := ValidateApiResponse(t, w)

	// Attempt to unlock using a good secret
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", fmt.Sprintf("?key=%v&secret=%v&duration=100", key, lockRes.Extras["secret"]), nil)
	g.HandleExtend(w, r)

	// Validate the response is successful
	ValidateApiResponse(t, w)
}