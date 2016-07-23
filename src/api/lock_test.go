package api

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/KyleBanks/glock/src/glock"
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

// TODO: Positive case