package cmdrun

import (
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
)

// RunHTTPOperation executes an HTTP operation and displays the result as a table (if any) for CLI commands
func RunHTTPOperation(op api.HTTPOperation, dryRun bool) {
	// If dry run, print operation as cURL command
	if dryRun {
		curlCmd, cErr := op.GetAsCurlCommand(true)
		if cErr != nil {
			color.Red(cErr.Error())
		}
		color.Green("%s", curlCmd)
		return
	}

	// Create spinner
	action, actionDone, name := op.GetDetails()
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp(action, actionDone, name)
	if sErr != nil {
		color.Red(sErr.Error())
		return
	}
	_ = spinner.Start()

	// Make HTTP request and get response
	tbl, rErr := op.GetResponseAsArray()

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
		return
	}

	_ = spinner.Stop()

	if tbl != nil {
		colNames := op.GetVerboseResponseFieldNames()
		utils.RenderDefaultTable(colNames, tbl)
	}

	return
}
