package api

import (
	"net/http"
	"github.com/KyleBanks/go-kit/router"
	"strconv"
	"github.com/KyleBanks/glock/src/glock"
)

const (
	LockRequired_Key = "key"
	LockRequired_Secret = "secret"

	LockOptional_Duration = "duration"
)


// HandleLock handles incoming API requests to perform a lock.
func (g glockApi) HandleLock(w http.ResponseWriter, r *http.Request) {
	// Validate required params
	if !router.HasParam(r, LockRequired_Key) {
		g.Write(NewFailedResponse(glock.ErrMissingKey), w, r)
		return
	} else if !router.HasParam(r, LockRequired_Secret) {
		g.Write(NewFailedResponse(glock.ErrMissingSecret), w, r)
		return
	}

	// Grab the required params ...
	key := router.Param(r, LockRequired_Key)
	secret := router.Param(r, LockOptional_Duration)

	// ... and the optional ones.
	var duration int
	if router.HasParam(r, LockOptional_Duration) {
		if d, err := strconv.Atoi(router.Param(r, LockOptional_Duration)); err != nil {
			g.Write(NewFailedResponse(glock.ErrInvalidDuration), w, r)
		} else {
			duration = d
		}
	}

	// Determine what to do based on if there's a duration
	var err *glock.GlockError
	if duration > 0 {
		err = g.glocker.LockWithDuration(key, secret, duration)
	} else {
		err = g.glocker.Lock(key, secret)
	}

	// Handle the response
	if err != nil {

		return
	}

	g.Write(SuccessResponse, w, r)
}