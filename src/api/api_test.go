package api

import (
	"github.com/KyleBanks/glock/src/log"
	"github.com/KyleBanks/glock/src/glock"
	"testing"
	"net/http/httptest"
	"encoding/json"
)

// NewForTest instantiates and returns a new glockApi instance for testing
// purposes.
func NewForTest() *glockApi {
	logger := log.New(false)
	return New(logger, 0, glock.New(logger))
}

// ValidateApiResponse checks if an API response follows
// valid formatting, headers, content-type, etc.
//
// If the response is not valid, the test will be halted with a fatal error
func ValidateApiResponse(t *testing.T, w *httptest.ResponseRecorder) *apiResponse {
	if w.HeaderMap.Get("Content-Type") != "application/json" {
		t.Fatalf("Bad content-type: %v", w.HeaderMap.Get("Content-Type"))
	} else if w.Code != 200 {
		t.Fatalf("Invalid response code [%v]: %v", w.Code, w.Body)
	}

	var res apiResponse
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatal(err)
	}

	if !res.Success {
		t.Fatalf("Unexpected unsuccessful response: %v", res)
	} else if res.Error != nil {
		t.Fatalf("Expected nil error to be returned for successful request: %v", res.Error)
	}

	return &res
}

// ValidateApiResponseError checks if an API response follows proper
// error formatting, headers, content-type, etc.
//
// If the response is not a valid error response, the test will be halted
// with a fatal error.
func ValidateApiResponseError(t *testing.T, w *httptest.ResponseRecorder) *apiResponse {
	if w.HeaderMap.Get("Content-Type") != "application/json" {
		t.Fatalf("Bad content-type: %v", w.HeaderMap.Get("Content-Type"))
	} else if w.Code != 400 {
		t.Fatalf("Invalid error response code [%v]: %v", w.Code, w.Body)
	}

	var res apiResponse
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatal(err)
	}

	if res.Success {
		t.Fatalf("Unexpected success response: %v", res)
	} else if res.Error == nil {
		t.Fatalf("Unexpected nil error returned: %v", res)
	}

	return &res
}
