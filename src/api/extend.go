package api

import (
	"net/http"
	"github.com/KyleBanks/go-kit/router"
	"github.com/KyleBanks/glock/src/glock"
	"strconv"
)

const (
	ExtendRequired_Key = LockRequired_Key
	ExtendRequired_Secret = UnlockRequired_Secret
	ExtendRequired_Duration = LockOptional_Duration
)

// HandleExtend handles incoming API requests to perform an extend.
func (g glockApi) HandleExtend(w http.ResponseWriter, r *http.Request) {
	// Validate required params
	if !router.HasParam(r, ExtendRequired_Key) {
		g.Write(NewFailedResponse(glock.ErrMissingKey), w, r)
		return
	} else if !router.HasParam(r, ExtendRequired_Secret) {
		g.Write(NewFailedResponse(glock.ErrMissingSecret), w, r)
		return
	} else if !router.HasParam(r, ExtendRequired_Duration) {
		g.Write(NewFailedResponse(glock.ErrMissingDuration), w, r)
		return
	}

	// Grad the required params
	key := router.Param(r, ExtendRequired_Key)
	secret := router.Param(r, ExtendRequired_Secret)
	var duration int
	if d, err := strconv.Atoi(router.Param(r, ExtendRequired_Duration)); err != nil {
		g.Write(NewFailedResponse(glock.ErrInvalidDuration), w, r)
		return
	} else {
		duration = d
	}

	// Perform the extension
	if err := g.glocker.Extend(key, secret, duration); err != nil {
		g.Write(NewFailedResponse(err), w, r)
		return
	}

	g.Write(NewSuccessResponse(map[string]string{}), w, r)
}
