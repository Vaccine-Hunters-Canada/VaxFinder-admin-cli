package update

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin requirement upate`
func CmdRunE(cmd *cobra.Command, args []string) error {

	id, aErr := strconv.Atoi(args[0])
	if aErr != nil {
		return errors.New("expecting id as integer")
	}

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
	body := api.UpdateRequirementApiV1RequirementsRequirementIdPutJSONRequestBody{Name: name, Description: description, Id: id}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("update", "updated", "requirement")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.UpdateRequirementApiV1RequirementsRequirementIdPutWithResponse(context.Background(), id, body)

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
