package utility

import "fmt"

type HttpError struct {
	Message string
	Status  int
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("status %d: %s", e.Status, e.Message)
}
