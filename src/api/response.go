package api

import (
	"fmt"
	"github.com/KyleBanks/glock/src/glock"
)

var (
	SuccessResponse = &apiResponse{success: true, error: nil}
)

type apiResponse struct {
	success bool `json:"success"`
	error *glock.GlockError `json:"error"`
}

// NewFailedResponse instantiates and returns a new apiResponse for a failure case.
func NewFailedResponse(err *glock.GlockError) *apiResponse {
	return &apiResponse{
		success: false,
		error: err,
	}
}

func (a apiResponse) String() string {
	return fmt.Sprintf(`{Success: %v, Error: %v}`, a.success, a.error)
}