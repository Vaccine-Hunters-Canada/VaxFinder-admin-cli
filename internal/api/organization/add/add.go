package add

import (
	"context"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var body api.CreateOrganizationApiV1OrganizationsPostJSONRequestBody
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
func (HTTPOperation) SetRequestBody(shortName string, fullName, description, url *string) error {
	body = api.CreateOrganizationApiV1OrganizationsPostJSONRequestBody{
		ShortName:   shortName,
		FullName:    fullName,
		Description: description,
		Url:         url,
	}
	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "add", "added", "organization"
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "short name", "full name", "description", "url", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.CreateOrganizationApiV1OrganizationsPostWithResponse(context.Background(), body)
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
				json.ShortName,
				utils.CoalesceString(json.FullName),
				utils.CoalesceString(json.Description),
				utils.CoalesceString(json.Url),
				json.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}
