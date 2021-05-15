package get

import (
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
	"vf-admin/internal/utils"
)

// CmdRunE is what's executed when running `vf-admin requirement get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id, aErr := strconv.Atoi(args[0])
	if aErr != nil {
		return errors.New("expecting integer as id")
	}

	// Create the API client using the authentication key for requests
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("get", "got", "requirement")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.RetrieveRequirementByIdApiV1RequirementsRequirementIdGetWithResponse(context.Background(), id)

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

	colNames := []string{"id", "name", "description", "created at"}
	data := [][]string{
		{
			strconv.Itoa(res.JSON200.Id), res.JSON200.Name, res.JSON200.Description, res.JSON200.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
