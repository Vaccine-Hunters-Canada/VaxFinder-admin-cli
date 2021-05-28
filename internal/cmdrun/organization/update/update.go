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

// CmdRunE is what's executed when running `vf-admin organization update`
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

	flags := cmd.Flags()
	shortName, _ := flags.GetString("shortName")
	var fullName, description, url *string
	if flags.Changed("fullName") {
		t, _ := flags.GetString("fullName")
		fullName = &t
	} else {
		fullName = nil
	}
	if flags.Changed("description") {
		t, _ := flags.GetString("description")
		description = &t
	} else {
		description = nil
	}
	if flags.Changed("url") {
		t, _ := flags.GetString("url")
		url = &t
	} else {
		url = nil
	}
	body := api.UpdateOrganizationApiV1OrganizationsOrganizationIdPutJSONRequestBody{ShortName: shortName, FullName: fullName, Description: description, Url: url}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("update", "updated", "organization")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.UpdateOrganizationApiV1OrganizationsOrganizationIdPutWithResponse(context.Background(), id, body)

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
