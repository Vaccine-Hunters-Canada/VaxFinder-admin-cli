package get

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin organization get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id, aErr := strconv.Atoi(args[0])
	if aErr != nil {
		return errors.New("expecting integer as id")
	}

	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("get", "got", "organization")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.RetrieveOrganizationByIdApiV1OrganizationsOrganizationIdGetWithResponse(context.Background(), id)

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
	json := res.JSON200
	data := [][]string{
		{
			strconv.Itoa(json.Id), json.ShortName, utils.CoalesceString(json.FullName), utils.CoalesceString(json.Description), utils.CoalesceString(json.Url), json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
