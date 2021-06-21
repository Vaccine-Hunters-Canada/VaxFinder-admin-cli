package list

import (
	"context"
	"vf-admin/internal/api"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var authKey string

var vaID string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	vaID = args[0]

	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "list", "got", "va timeslot"
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

	res, rErr := client.ListTimeslotsForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsGetWithResponse(context.Background(), vaID)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		var data [][]string
		for _, row := range *res.JSON200 {
			var takenAt string
			if row.TakenAt == nil {
				takenAt = ""
			} else {
				takenAt = row.TakenAt.String()
			}
			data = append(data, []string{
				row.Id, row.VaccineAvailability, row.Time.String(), takenAt, row.CreatedAt.String(),
			})
		}
		return data, nil
	}

	return nil, nil
}
