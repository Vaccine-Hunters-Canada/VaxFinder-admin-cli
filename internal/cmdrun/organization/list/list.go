package list

import (
	"context"
	"strconv"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin organization list`
func CmdRunE(cmd *cobra.Command, args []string) error {
	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "organizations")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.ListOrganizationsApiV1OrganizationsGetWithResponse(context.Background())

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

	colNames := []string{"id", "short name", "full name", "description", "url", "created at"}

	var data [][]string
	for _, row := range *res.JSON200 {
		data = append(data, []string{
			strconv.Itoa(row.Id), row.ShortName, utils.CoalesceString(row.FullName), utils.CoalesceString(row.Description), utils.CoalesceString(row.Url), row.CreatedAt.String(),
		})
	}

	utils.RenderDefaultTable(colNames, data)

	return nil
}
