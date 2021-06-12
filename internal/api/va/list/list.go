package list

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var authKey string
var id string
var postCode string
var minDate *openapi_types.Date

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	return nil
}

// SetRequestURLArgumentsFromFlags sets the appropriate url arguments from the command flags
func (HTTPOperation) SetRequestURLArgumentsFromFlags(flags *pflag.FlagSet) error {
	postCode, _ = flags.GetString("postcode")

	if flags.Changed("mindate") {
		t, _ := flags.GetString("mindate")
		t2, tErr := time.Parse("2006-01-02", t)
		if tErr != nil {
			color.Red(tErr.Error())
			return nil
		}
		minDate = &openapi_types.Date{Time: t2}
	}

	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "list", "got", "vaccine availability"
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "organization", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	params := api.ListVaccineAvailabilityApiV1VaccineAvailabilityGetParams{PostalCode: postCode, MinDate: minDate}

	res, rErr := client.ListVaccineAvailabilityApiV1VaccineAvailabilityGetWithResponse(context.Background(), &params)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		var data [][]string
		for _, row := range *res.JSON200 {
			var locationLine1, locationLine2, locationProvince, org *string
			if row.Location.Address != nil {
				locationLine1 = row.Location.Address.Line1
				locationLine2 = row.Location.Address.Line2
				locationProvince = &row.Location.Address.Province
				if row.Location.Organization != nil {
					org = row.Location.Organization.FullName
				}
			}

			data = append(data, []string{
				row.Id,
				row.Date.String(),
				strconv.Itoa(row.NumberAvailable),
				utils.CoalesceInt(row.NumberTotal),
				utils.CoalesceInt(row.Vaccine),
				strconv.Itoa(int(row.InputType)),
				utils.CoalesceString(row.Tags),
				fmt.Sprintf("%d - %s %s %s %s", row.Location.Id, row.Location.Name, utils.CoalesceString(locationLine1), utils.CoalesceString(locationLine2), utils.CoalesceString(locationProvince)),
				utils.CoalesceString(org),
				row.CreatedAt.String(),
			})
		}
		return data, nil
	}

	return nil, nil
}
