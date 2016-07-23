package glock

import (
	"testing"
	"strings"
	"strconv"
)

func TestGlockError_String(t *testing.T) {
	s := ErrInvalidDuration.String()

	if !strings.Contains(s, ErrInvalidDuration.Message) {
		t.Fatalf("Invalid String() response, missing message: %v", s)
	} else if !strings.Contains(s, strconv.Itoa(ErrInvalidDuration.Code)) {
		t.Fatalf("Invalid String() response, missing code: %v", s)
	}
}
