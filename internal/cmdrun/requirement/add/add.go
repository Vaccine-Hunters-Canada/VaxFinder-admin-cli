package add

import (
	"context"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// CmdRunE is what's executed when running `vf-admin requirement add`
func CmdRunE(cmd *cobra.Command, args []string) error {
	// Retrieve the authentication key from configuration file
	key, kErr := utils.GetKeyFromProfile(cmd)
	if kErr != nil {
		color.Red(kErr.Error())
		return nil
	}

	// Create the API client using the authentication key for requests
	client, cErr := api.GetAPIClientFromKey(key)
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	body := api.CreateRequirementApiV1RequirementsPostJSONRequestBody{Name: name, Description: description}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("add", "added", "requirement")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.CreateRequirementApiV1RequirementsPostWithResponse(context.Background(), body)

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
