package repository

import (
	"fmt"
	"net/http"
)

const (
	userErrorHundreds     = 4
	serverErrorHundreds   = 5
	httpErrorCategorySize = 100
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode/httpErrorCategorySize == userErrorHundreds:
		// errors 4xx
		return repositoryError(fmt.Sprintf("user-end error: %d", statusCode))
	case statusCode/httpErrorCategorySize == serverErrorHundreds:
		// errors 5xx
		return repositoryError(fmt.Sprintf("server error: %d", statusCode))
	case statusCode != http.StatusOK:
		// any other usecases
		return repositoryError(fmt.Sprintf("invalid status code: %d", statusCode))
	}
	return nil
}
