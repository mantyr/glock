package api

import "fmt"

var (
	ErrMissingKey = apiError{Code: 1, Message: "Missing required param: 'key'"}
	ErrMissingSecret = apiError{Code: 2, Message: "Missing required param: 'secret'"}
)

type apiError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (a apiError) String() string {
	return fmt.Sprintf(`{code: %v, message: %v}`, a.Code, a.Message)
}