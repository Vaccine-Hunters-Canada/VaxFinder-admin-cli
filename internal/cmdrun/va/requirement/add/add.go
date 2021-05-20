package add

import (
	"context"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va requirement add <va-id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id := args[0]

	// Retrieve the authentication key from configuration file
	key, kErr := utils.GetKeyFromProfile(cmd)
	if kErr != nil {
		color.Red(kErr.Error())
		return nil
	}

	// Create the API client using the authentication key for requests
	client, cErr := utils.GetAPIClientFromKey(key)
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	flags := cmd.Flags()
	requirement, _ := flags.GetInt("requirement")
	body := api.CreateRequirementForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdRequirementsPostJSONRequestBody{Requirement: requirement}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("add", "added", "va requirement")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.CreateRequirementForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdRequirementsPostWithResponse(context.Background(), id, body)

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

	colNames := []string{"id", "vaccine availability", "requirement", "active", "name", "description", "created at"}
	json := res.JSON200
	data := [][]string{
		{
			json.Id, json.VaccineAvailability, strconv.Itoa(json.Requirement), strconv.FormatBool(json.Active), json.Name, json.Description, json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}