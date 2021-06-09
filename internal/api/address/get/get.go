package get

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
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
	return "get", "got", "address " + strconv.Itoa(id)
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "line 1", "line 2", "city", "postal code", "province" /* Coming soon: , "latitude", "longitude"*/, "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.RetrieveAddressByIdApiV1AddressesAddressIdGetWithResponse(context.Background(), id)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		json := res.JSON200
		return [][]string{
			{
				strconv.Itoa(json.Id), utils.CoalesceString(json.Line1), utils.CoalesceString(json.Line2), utils.CoalesceString(json.City), json.Postcode, json.Province /*json.Latitude.String(), json.Longitude.String(),*/, json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}
