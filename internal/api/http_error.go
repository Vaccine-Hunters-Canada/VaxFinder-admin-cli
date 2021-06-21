package api

import (
	"errors"
	"net/http"
)

// HandleHTTPError returns an error based on the status code
func HandleHTTPError(code int, body []byte) error {
	// TODO: Print error messages differently depending on status code
	return errors.New(http.StatusText(code) + ": " + string(body))
}
