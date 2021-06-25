package add

import (
	"context"
	"moul.io/http2curl"
	"strconv"
	"time"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var body api.CreateVaccineAvailabilityApiV1VaccineAvailabilityPostJSONRequestBody
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	return nil
}

// SetRequestBody sets the appropriate body for the HTTP operation
func (HTTPOperation) SetRequestBody(time time.Time, inputType api.InputTypeEnum, location int, numberAvailable int, numberTotal *int, tags *string, vaccine *int) error {
	body = api.CreateVaccineAvailabilityApiV1VaccineAvailabilityPostJSONRequestBody{
		Date:            time,
		InputType:       inputType,
		Location:        location,
		NumberAvailable: numberAvailable,
		NumberTotal:     numberTotal,
		Tags:            tags,
		Vaccine:         vaccine,
	}
	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "add", "added", "vaccine availability"
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.CreateVaccineAvailabilityApiV1VaccineAvailabilityPostWithResponse(context.Background(), body)
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
				json.Id,
				json.Date.String(),
				strconv.Itoa(json.NumberAvailable),
				utils.CoalesceInt(json.NumberTotal),
				utils.CoalesceInt(json.Vaccine),
				strconv.Itoa(int(json.InputType)),
				utils.CoalesceString(json.Tags),
				strconv.Itoa(json.Location),
				json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}

// GetAsCurlCommand returns the HTTP operation as a cURL command
func (HTTPOperation) GetAsCurlCommand(withKey bool) (*http2curl.CurlCommand, error) {
	// Create the HTTP Request (struct)
	req, rErr := api.NewCreateVaccineAvailabilityApiV1VaccineAvailabilityPostRequest(utils.GetBaseURL(), body)
	if rErr != nil {
		return nil, rErr
	}
	// Attach auth key to request if it exists
	if authKey != "" && withKey {
		req.Header.Set("Authorization", "Bearer "+authKey)
	}
	return http2curl.GetCurlCommand(req)
}
