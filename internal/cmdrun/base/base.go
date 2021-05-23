package base

import (
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Tabler is used to convert data to string tables
type Tabler interface {
	GetHeaders() []string
	GetResourceName() (string, string, string)
	GetArray(cmd *cobra.Command, args []string, client *api.ClientWithResponses) ([][]string, error)
}

// Cmd executes the request and displays the result or errors as a table (if any)
func Cmd(cmd *cobra.Command, args []string, tabler Tabler) error {
	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	action, actionDone, name := tabler.GetResourceName()
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp(action, actionDone, name)
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	tbl, rErr := tabler.GetArray(cmd, args, client)

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
		return nil
	}

	_ = spinner.Stop()

	if tbl != nil {
		colNames := tabler.GetHeaders()
		utils.RenderDefaultTable(colNames, tbl)
	}

	return nil
}
