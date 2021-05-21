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
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va list`
func CmdRunE(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	postcode, _ := flags.GetString("postcode")

	var minDate *openapi_types.Date

	if flags.Changed("mindate") {
		t, _ := flags.GetString("mindate")
		t2, tErr := time.Parse("2006-01-02", t)
		if tErr != nil {
			color.Red(tErr.Error())
			return nil
		}
		minDate = &openapi_types.Date{Time: t2}
	}

	params := api.ListVaccineAvailabilityApiV1VaccineAvailabilityGetParams{PostalCode: postcode, MinDate: minDate}

	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "va")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.ListVaccineAvailabilityApiV1VaccineAvailabilityGetWithResponse(context.Background(), &params)

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
		return nil
	}

	if res.StatusCode() != 200 {
		spinner.StopFailMessage(res.Status() + ": " + string(res.Body))
		_ = spinner.StopFail()
		return nil
	}

	_ = spinner.Stop()

	colNames := []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "organization", "created at"}
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

	utils.RenderDefaultTable(colNames, data)

	return nil
}
