package list

import (
	"context"
	"strconv"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin address list`
func CmdRunE(cmd *cobra.Command, args []string) error {
	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "address")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.ListAddressesApiV1AddressesGetWithResponse(context.Background())

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

	colNames := []string{"id", "line 1", "line 2", "city", "postal code", "province" /* Coming soon: , "latitude", "longitude"*/, "created at"}
	var data [][]string
	for _, row := range *res.JSON200 {
		data = append(data, []string{
			strconv.Itoa(row.Id), utils.CoalesceString(row.Line1), utils.CoalesceString(row.Line2), utils.CoalesceString(row.City), row.Postcode, row.Province /*row.Latitude.String(), row.Longitude.String(),*/, row.CreatedAt.String(),
		})
	}

	utils.RenderDefaultTable(colNames, data)

	return nil
}
