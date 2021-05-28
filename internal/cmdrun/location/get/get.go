package get

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
	locationutils "vf-admin/internal/cmdrun/location"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin location get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id, aErr := strconv.Atoi(args[0])
	if aErr != nil {
		return errors.New("expecting id as integer")
	}

	// Create the API client
	client, cErr := api.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("get", "got", "location")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.RetrieveLocationByIdApiV1LocationsLocationIdGetWithResponse(context.Background(), id)

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

	colNames := []string{"id", "name", "active", "postcode", "phone", "notes", "url", "tags", "org", "address", "created at"}
	json := res.JSON200
	data := [][]string{
		locationutils.ConvertExpandedLocationJSONToTableRow(json),
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
