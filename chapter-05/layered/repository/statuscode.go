package repository

import (
	"fmt"
	"net/http"
)

func checkStatusCode(statusCode int) error {
	// TODO: Handle 4xx and 5xx separately ?
	if statusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", statusCode)
	}
	return nil
}
