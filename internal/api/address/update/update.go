package update

import (
	"context"
	"errors"
	"moul.io/http2curl"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var body = api.UpdateAddressApiV1AddressesAddressIdPutJSONRequestBody{}
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
	body.Id = converted

	return nil
}

// SetRequestBody sets the appropriate body for the HTTP operation
func (HTTPOperation) SetRequestBody(province, postcode string, latitude, longitude float32, line1, line2, city *string) error {
	body.Province = province
	body.Postcode = postcode
	body.Latitude = latitude
	body.Longitude = longitude
	body.Line1 = line1
	body.Line2 = line2
	body.City = city
	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "update", "updated", "address " + strconv.Itoa(body.Id)
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

	res, rErr := client.UpdateAddressApiV1AddressesAddressIdPutWithResponse(context.Background(), body.Id, body)
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
				strconv.Itoa(json.Id),
				utils.CoalesceString(json.Line1),
				utils.CoalesceString(json.Line2),
				utils.CoalesceString(json.City),
				json.Postcode,
				json.Province /*json.Latitude.String(), json.Longitude.String(),*/,
				utils.GetFromNow(json.CreatedAt),
			},
		}, nil
	}

	return nil, nil
}

// GetAsCurlCommand returns the HTTP operation as a cURL command
func (HTTPOperation) GetAsCurlCommand(withKey bool) (*http2curl.CurlCommand, error) {
	// Create the HTTP Request (struct)
	req, rErr := api.NewUpdateAddressApiV1AddressesAddressIdPutRequest(utils.GetBaseURL(), body.Id, body)
	if rErr != nil {
		return nil, rErr
	}
	// Attach auth key to request if it exists
	if authKey != "" && withKey {
		req.Header.Set("Authorization", "Bearer "+authKey)
	}
	return http2curl.GetCurlCommand(req)
}
