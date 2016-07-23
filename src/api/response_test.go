package api

import (
	"testing"
	"github.com/KyleBanks/glock/src/glock"
	"time"
	"reflect"
)

func TestNewFailedResponse(t *testing.T) {
	res := NewFailedResponse(glock.ErrMissingSecret)

	if res.Success {
		t.Fatalf("Expected Success to be false: %v", res)
	} else if res.Error == nil {
		t.Fatalf("Unexpected nil Error returned: %v", res)
	} else if res.Error.Code != glock.ErrMissingSecret.Code || res.Error.Message != glock.ErrMissingSecret.Message {
		t.Fatalf("Unexpected value returned for Error: %v", res.Error)
	}
}

func TestNewSuccessResponse(t *testing.T) {
	extras := map[string]string {
		"testing": time.Now().String(),
	}

	res := NewSuccessResponse(extras)

	if !res.Success {
		t.Fatalf("Expected Success to be true: %v", res)
	} else if res.Error != nil {
		t.Fatalf("Expected Error to be nil: %v", res)
	} else if !reflect.DeepEqual(res.Extras, extras) {
		t.Fatalf("Unexpected extras returned: %v", res)
	}
}