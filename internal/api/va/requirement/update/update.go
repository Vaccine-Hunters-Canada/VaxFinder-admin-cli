package update

import (
	"context"
	"strconv"
	"vf-admin/internal/api"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var vaID string
var reqID string
var body = api.UpdateRequirementForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdRequirementsRequirementIdPutJSONRequestBody{}
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	vaID = args[0]
	reqID = args[1]

	return nil
}

// SetRequestBody sets the appropriate body for the HTTP operation
func (HTTPOperation) SetRequestBody(requirement int, active bool) error {
	body.Requirement = requirement
	body.Active = active
	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "update", "updated", "va requirement " + reqID
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "vaccine availability", "requirement", "active", "name", "description", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.UpdateRequirementForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdRequirementsRequirementIdPutWithResponse(context.Background(), vaID, reqID, body)
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
				json.Id, json.VaccineAvailability, strconv.Itoa(json.Requirement), strconv.FormatBool(json.Active), json.Name, json.Description, json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}
