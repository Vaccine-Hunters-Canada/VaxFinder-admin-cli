package get

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var id int
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	converted, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("expecting id as integer")
	}
	id = converted

	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "get", "got", "requirement " + strconv.Itoa(id)
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "name", "description", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.RetrieveRequirementByIdApiV1RequirementsRequirementIdGetWithResponse(context.Background(), id)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		return [][]string{
			{
				strconv.Itoa(res.JSON200.Id), res.JSON200.Name, res.JSON200.Description, res.JSON200.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}