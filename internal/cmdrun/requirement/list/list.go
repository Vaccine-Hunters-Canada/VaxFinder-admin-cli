package list

import (
	"context"
	"strconv"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin requirement list`
func CmdRun(cmd *cobra.Command, args []string) {

	// Create the API client using the authentication key for requests
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "requirements")
	if sErr != nil {
		color.Red(sErr.Error())
	}
	_ = spinner.Start()

	res, rErr := client.ListRequirementsApiV1RequirementsGetWithResponse(context.Background())

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
	}

	if res.StatusCode() != 200 {
		spinner.StopFailMessage(res.Status() + ": " + string(res.Body))
		_ = spinner.StopFail()
	}

	_ = spinner.Stop()

	colNames := []string{"id", "name", "description", "created at"}
	var data [][]string

	for _, row := range *res.JSON200 {
		data = append(data, []string{
			strconv.Itoa(row.Id), row.Name, row.Description, row.CreatedAt.String(),
		})
	}

	utils.RenderDefaultTable(colNames, data)

}
