package api

import (
	"fmt"
	"github.com/KyleBanks/glock/src/glock"
)

var (
	SuccessResponse = &apiResponse{Success: true, Error: nil}
)

type apiResponse struct {
	Success bool `json:"success"`
	Error   *glock.GlockError `json:"error"`

	Extras map[string]string `json:"extras"`
}

// NewFailedResponse instantiates and returns a new apiResponse for a failure case.
func NewFailedResponse(err *glock.GlockError) *apiResponse {
	return &apiResponse{
		Success: false,
		Error: err,
	}
}

// NewSuccessRepsonse initializes and returns a success apiResponse with a map of extras.
func NewSuccessResponse(extras map[string]string) *apiResponse {
	return &apiResponse{
		Success: true,
		Extras: extras,
	}
}

func (a apiResponse) String() string {
	return fmt.Sprintf(`{Success: %v, Error: %v}`, a.Success, a.Error)
}