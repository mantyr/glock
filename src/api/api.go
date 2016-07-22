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
	Logger *log.Logger

	Port int

	Glocker glock.Glocker
}

// New instantiates and returns a glockApi instance.
func New(logger *log.Logger, port int, glocker glock.Glocker) *glockApi {
	return &glockApi {
		Logger: logger,

		Port: port,

		Glocker: glocker,
	}
}

// Run starts the glockApi server.
func (g glockApi) Run() {
	http.HandleFunc("/lock", g.lockHandler)

	g.Logger.Error(http.ListenAndServe(fmt.Sprintf(":%v", g.Port), nil))
}

// Write outputs the ApiResponse as JSON to the ResponseWriter.
func (g glockApi) Write(a *apiResponse, w http.ResponseWriter, r *http.Request) {
	g.Logger.Printf("%v: %v", r.RequestURI, a)
	if a.Success {
		w.WriteHeader(codeSuccess)
	} else {
		w.WriteHeader(codeError)
	}

	w.Header().Add(contentTypeHeaderName, contentTypeHeaderValue)
	json.NewEncoder(w).Encode(a)
}