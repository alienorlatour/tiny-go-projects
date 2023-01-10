package repository

import (
	"fmt"
	"net/http"
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		// errors 4xx
		return repositoryError(fmt.Sprintf("user-end error: %d", statusCode))
	case statusCode >= http.StatusInternalServerError:
		// errors 5xx
		return repositoryError(fmt.Sprintf("server error: %d", statusCode))
	case statusCode != http.StatusOK:
		// any other errors
		return repositoryError(fmt.Sprintf("invalid status code: %d", statusCode))
	}
	return nil
}
