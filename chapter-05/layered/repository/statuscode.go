package repository

import (
	"fmt"
	"net/http"
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode >= http.StatusBadRequest:
		return fmt.Errorf("error from the user %d", statusCode)
	case statusCode >= http.StatusInternalServerError:
		return fmt.Errorf("error from the server %d", statusCode)
	case statusCode != http.StatusOK:
		return fmt.Errorf("invalid status code %d", statusCode)
	}
	return nil
}
