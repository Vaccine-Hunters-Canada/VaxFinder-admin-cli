package get

import (
	"context"
	"fmt"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id := args[0]

	// Create the API client
	client, cErr := api.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("get", "got", "va")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.RetrieveVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdGetWithResponse(context.Background(), id)

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

	colNames := []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "organization", "created at"}
	json := res.JSON200
	var locationLine1, locationLine2, locationProvince, org *string
	if json.Location.Address != nil {
		locationLine1 = json.Location.Address.Line1
		locationLine2 = json.Location.Address.Line2
		locationProvince = &json.Location.Address.Province
		if json.Location.Organization != nil {
			org = json.Location.Organization.FullName
		}
	}
	data := [][]string{
		{
			json.Id,
			json.Date.String(),
			strconv.Itoa(json.NumberAvailable),
			utils.CoalesceInt(json.NumberTotal),
			utils.CoalesceInt(json.Vaccine),
			strconv.Itoa(int(json.InputType)),
			utils.CoalesceString(json.Tags),
			strconv.Itoa(json.Location.Id),
			fmt.Sprintf("%d - %s %s %s %s", json.Location.Id, json.Location.Name, utils.CoalesceString(locationLine1), utils.CoalesceString(locationLine2), utils.CoalesceString(locationProvince)),
			utils.CoalesceString(org),
			json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
