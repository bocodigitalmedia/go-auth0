package auth0mgmt

import "fmt"

type ServiceError struct {
	Name        string `json:"error"`
	Description string `json:"error_description"`
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Description)
}

func (e *ServiceError) IsEmpty() bool {
	return len(e.Name) > 0
}
