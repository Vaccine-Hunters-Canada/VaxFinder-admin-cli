package update

import (
	"context"
	"time"
	"vf-admin/internal/api"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var vaID string
var tsID string
var body = api.UpdateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsTimeslotIdPutJSONRequestBody{}
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	vaID = args[0]
	tsID = args[1]

	return nil
}

// SetRequestBody sets the appropriate body for the HTTP operation
func (HTTPOperation) SetRequestBody(time time.Time, takenAt *interface{}) error {
	body.Time = time
	body.TakenAt = takenAt

	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "update", "updated", "va timeslot " + tsID
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "vaccine availability", "time", "taken at", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.UpdateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsTimeslotIdPutWithResponse(context.Background(), vaID, tsID, body)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		json := res.JSON200
		var takenAt string
		if json.TakenAt == nil {
			takenAt = ""
		} else {
			takenAt = json.TakenAt.String()
		}
		return [][]string{
			{
				json.Id, json.VaccineAvailability, json.Time.String(), takenAt, json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}
