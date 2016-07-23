package api

import (
	"net/http"
	"github.com/KyleBanks/go-kit/router"
	"github.com/KyleBanks/glock/src/glock"
)

const (
	UnlockRequired_Key = LockRequired_Key
	UnlockRequired_Secret = "secret"
)

// HandleUnlock handles incoming API requests to perform an unlock.
func (g glockApi) HandleUnlock(w http.ResponseWriter, r *http.Request) {
	// Validate required params
	if !router.HasParam(r, UnlockRequired_Key) {
		g.Write(NewFailedResponse(glock.ErrMissingKey), w, r)
		return
	} else if !router.HasParam(r, UnlockRequired_Secret) {
		g.Write(NewFailedResponse(glock.ErrMissingSecret), w, r)
		return
	}

	// Grab the required params
	key := router.Param(r, UnlockRequired_Key)
	secret := router.Param(r, UnlockRequired_Secret)

	// Perform the unlock
	if err := g.glocker.Unlock(key, secret); err != nil {
		g.Write(NewFailedResponse(err), w, r)
		return
	}

	g.Write(NewSuccessResponse(map[string]string{}), w, r)
}