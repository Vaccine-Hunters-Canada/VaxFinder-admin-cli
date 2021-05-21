package list

import (
	"context"
	locationutils "vf-admin/internal/cmdrun/location"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRun is what's executed when running `vf-admin location list`
func CmdRun(cmd *cobra.Command, args []string) {
	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "locations")
	if sErr != nil {
		color.Red(sErr.Error())
	}
	_ = spinner.Start()

	res, rErr := client.ListLocationsApiV1LocationsGetWithResponse(context.Background())

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
	}

	if res.StatusCode() != 200 {
		spinner.StopFailMessage(res.Status() + ": " + string(res.Body))
		_ = spinner.StopFail()
	}

	_ = spinner.Stop()

	colNames := []string{"id", "name", "active", "postcode", "phone", "notes", "url", "tags", "org", "address", "created at"}

	var data [][]string
	for _, row := range *res.JSON200 {
		data = append(data, locationutils.ConvertExpandedLocationJSONToTableRow(&row))
	}

	utils.RenderDefaultTable(colNames, data)

}
