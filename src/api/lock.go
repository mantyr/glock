package api

import (
	"net/http"
	"github.com/KyleBanks/go-kit/router"
)

const (
	LockRequired_Key = "key"
	LockRequired_Secret = "secret"

	LockOptional_Duration = "duration"
)


// lockHandler handles incoming API requests to perform a lock.
func (g glockApi) lockHandler(w http.ResponseWriter, r *http.Request) {
	if !router.HasParam(r, LockRequired_Key) {
		g.Write(NewFailedResponse(ErrMissingKey), w, r)
		return
	} else if !router.HasParam(r, LockRequired_Secret) {
		g.Write(NewFailedResponse(ErrMissingSecret), w, r)
		return
	}

	g.Glocker.Lock("test", "secret")
}