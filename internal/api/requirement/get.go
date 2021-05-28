package requirement

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
)

// Get abstracts away the HTTP GET operation for requirement
type Get struct{}

var id int

// GetDetails returns the details of the HTTP operation
func (Get) GetDetails() (string, string, string) {
	return "get", "got", "requirement"
}

// SetRequestURLArguments sets the appropriate arguments for the HTTP operation
func (Get) SetRequestURLArguments(args []string) error {
	converted, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("expecting id as integer")
	}
	id = converted

	return nil
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (Get) GetVerboseResponseFieldNames() []string {
	return []string{"id", "name", "description", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (Get) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClient()
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
