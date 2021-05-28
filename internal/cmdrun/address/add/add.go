package add

import (
	"context"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin address add`
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

	flags := cmd.Flags()
	province, _ := flags.GetString("province")
	postcode, _ := flags.GetString("postcode")
	latitude, _ := flags.GetFloat32("latitude")
	longitude, _ := flags.GetFloat32("longitude")
	var line1, line2, city *string
	if flags.Changed("line1") {
		t, _ := flags.GetString("line1")
		line1 = &t
	} else {
		line1 = nil
	}
	if flags.Changed("line2") {
		t, _ := flags.GetString("line2")
		line2 = &t
	} else {
		line2 = nil
	}
	if flags.Changed("city") {
		t, _ := flags.GetString("city")
		city = &t
	} else {
		city = nil
	}

	body := api.CreateAddressApiV1AddressesPostJSONRequestBody{Province: province, Postcode: postcode, Latitude: latitude, Longitude: longitude, Line1: line1, Line2: line2, City: city}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("add", "added", "address")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.CreateAddressApiV1AddressesPostWithResponse(context.Background(), body)

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

	colNames := []string{"id", "line 1", "line 2", "city", "postal code", "province" /* Coming soon: , "latitude", "longitude"*/, "created at"}
	json := res.JSON200
	data := [][]string{
		{
			strconv.Itoa(json.Id), utils.CoalesceString(json.Line1), utils.CoalesceString(json.Line2), utils.CoalesceString(json.City), json.Postcode, json.Province /*json.Latitude.String(), json.Longitude.String(),*/, json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
