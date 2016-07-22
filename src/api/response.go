package api

import "fmt"

type apiResponse struct {
	Success bool `json:"success"`
	Error *apiError `json:"error"`
}

// NewFailedResponse instantiates and returns a new apiResponse for a failure case.
func NewFailedResponse(err apiError) *apiResponse {
	return &apiResponse{
		Success: false,
		Error: &err,
	}
}

func (a apiResponse) String() string {
	return fmt.Sprintf(`{Success: %v, Error: %v}`, a.Success, a.Error)
}