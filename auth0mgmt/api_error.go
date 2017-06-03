package auth0mgmt

import "fmt"

type ApiError struct {
	StatusCode int64  `json:"statusCode,omitempty"`
	Code       string `json:"errorCode,omitempty"`
	Name       string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d %s (%s): %s", e.StatusCode, e.Code, e.Name, e.Message)
}

func (e *ApiError) IsEmpty() bool {
	return e.StatusCode == 0
}
