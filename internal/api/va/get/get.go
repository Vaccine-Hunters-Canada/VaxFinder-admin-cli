package get

import (
	"context"
	"fmt"
	"moul.io/http2curl"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
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
	return "get", "got", "vaccine availability " + id
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "location", "organization", "created at"}

}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.RetrieveVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdGetWithResponse(context.Background(), id)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		json := res.JSON200

		var locationLine1, locationLine2, locationProvince, org *string
		if json.Location.Address != nil {
			locationLine1 = json.Location.Address.Line1
			locationLine2 = json.Location.Address.Line2
			locationProvince = &json.Location.Address.Province
			if json.Location.Organization != nil {
				org = json.Location.Organization.FullName
			}
		}

		return [][]string{
			{
				json.Id,
				json.Date.String(),
				strconv.Itoa(json.NumberAvailable),
				utils.CoalesceInt(json.NumberTotal),
				utils.CoalesceInt(json.Vaccine),
				strconv.Itoa(int(json.InputType)),
				utils.CoalesceString(json.Tags),
				strconv.Itoa(json.Location.Id),
				fmt.Sprintf("%d - %s %s %s %s", json.Location.Id, json.Location.Name, utils.CoalesceString(locationLine1), utils.CoalesceString(locationLine2), utils.CoalesceString(locationProvince)),
				utils.CoalesceString(org),
				json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}

// GetAsCurlCommand returns the HTTP operation as a cURL command
func (HTTPOperation) GetAsCurlCommand(withKey bool) (*http2curl.CurlCommand, error) {
	// Create the HTTP Request (struct)
	req, rErr := api.NewRetrieveVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdGetRequest(utils.GetBaseURL(), id)
	if rErr != nil {
		return nil, rErr
	}
	// Attach auth key to request if it exists
	if authKey != "" && withKey {
		req.Header.Set("Authorization", "Bearer "+authKey)
	}
	return http2curl.GetCurlCommand(req)
}
