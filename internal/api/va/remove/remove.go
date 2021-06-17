package remove

import (
	"context"
	"vf-admin/internal/api"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var id string
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	id = args[0]

	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "remove", "removed", "vaccine availability " + id
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return nil
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.DeleteVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdDeleteWithResponse(context.Background(), id)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 204 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	return nil, nil
}