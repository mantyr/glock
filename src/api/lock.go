package api

import (
	"net/http"
	"github.com/KyleBanks/go-kit/router"
	"strconv"
	"github.com/KyleBanks/glock/src/glock"
)

const (
	LockRequired_Key = "key"
	LockOptional_Duration = "duration"

	ExtraSecret = "secret"
)


// HandleLock handles incoming API requests to perform a lock.
func (g glockApi) HandleLock(w http.ResponseWriter, r *http.Request) {
	// Validate required params
	if !router.HasParam(r, LockRequired_Key) {
		g.Write(NewFailedResponse(glock.ErrMissingKey), w, r)
		return
	}

	// Grab the required params ...
	key := router.Param(r, LockRequired_Key)

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
	var secret string
	if duration > 0 {
		secret, err = g.glocker.LockWithDuration(key, duration)
	} else {
		secret, err = g.glocker.Lock(key)
	}

	// Handle the response
	if err != nil {
		g.Write(NewFailedResponse(err), w, r)
		return
	}

	g.Write(NewSuccessResponse(map[string]string{
		ExtraSecret: secret,
	}), w, r)
}