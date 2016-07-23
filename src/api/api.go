// Package "api" handles the REST API services. It manages listening for and dispatching incoming requests
// to the underlying glock store.
package api

import (
	"net/http"
	"fmt"
	"github.com/KyleBanks/glock/src/log"
	"github.com/KyleBanks/glock/src/glock"
	"encoding/json"
)

const (
	// Http response code for successful requests.
	codeSuccess = 200

	// Http response code for failed requests.
	codeError = 400

	// Response content type header
	contentTypeHeaderName = "Content-Type"

	// Response content type value
	contentTypeHeaderValue = "application/json"
)

type glockApi struct {
	logger *log.Logger

	port int

	glocker glock.Glocker
}

// New instantiates and returns a glockApi instance.
func New(logger *log.Logger, port int, glocker glock.Glocker) *glockApi {
	return &glockApi {
		logger: logger,

		port: port,

		glocker: glocker,
	}
}

// Run starts the glockApi server.
func (g glockApi) Run() {
	http.HandleFunc("/lock", g.HandleLock)

	g.logger.Error(http.ListenAndServe(fmt.Sprintf(":%v", g.port), nil))
}

// Write outputs the ApiResponse as JSON to the ResponseWriter.
func (g glockApi) Write(a *apiResponse, w http.ResponseWriter, r *http.Request) {
	g.logger.Printf("%v: %v", r.RequestURI, a)

	// Set the response code
	if a.Success {
		w.WriteHeader(codeSuccess)
	} else {
		w.WriteHeader(codeError)
	}

	// Set the content-type
	w.Header().Add(contentTypeHeaderName, contentTypeHeaderValue)

	// Write the response
	if err := json.NewEncoder(w).Encode(a); err != nil {
		g.logger.Error(err)
	}
}