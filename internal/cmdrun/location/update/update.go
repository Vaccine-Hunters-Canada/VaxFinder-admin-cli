package update

import (
	"context"
	"errors"
	"strconv"
	"vf-admin/internal/api"
	locationutils "vf-admin/internal/cmdrun/location"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin location update`
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
	client, cErr := utils.GetAPIClientFromKey(key)
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	flags := cmd.Flags()
	active, _ := flags.GetInt("active")
	name, _ := flags.GetString("name")
	var postcode, phone, notes, url, tags *string
	var org, address *int
	if flags.Changed("postcode") {
		t, _ := flags.GetString("postcode")
		postcode = &t
	} else {
		postcode = nil
	}
	if flags.Changed("phone") {
		t, _ := flags.GetString("phone")
		phone = &t
	} else {
		phone = nil
	}
	if flags.Changed("notes") {
		t, _ := flags.GetString("notes")
		notes = &t
	} else {
		notes = nil
	}
	if flags.Changed("url") {
		t, _ := flags.GetString("url")
		url = &t
	} else {
		url = nil
	}
	if flags.Changed("tags") {
		t, _ := flags.GetString("tags")
		tags = &t
	} else {
		tags = nil
	}
	if flags.Changed("organization") {
		t, _ := flags.GetInt("organization")
		org = &t
	} else {
		org = nil
	}
	if flags.Changed("address") {
		t, _ := flags.GetInt("address")
		address = &t
	} else {
		address = nil
	}

	body := api.UpdateLocationApiV1LocationsLocationIdPutJSONRequestBody{Active: active, Name: name, Postcode: postcode, Phone: phone, Notes: notes, Url: url, Tags: tags, Organization: org, Address: address}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("update", "updated", "location")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.UpdateLocationApiV1LocationsLocationIdPutWithResponse(context.Background(), id, body)

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

	colNames := []string{"id", "name", "active", "postcode", "phone", "notes", "url", "tags", "org", "address", "created at"}
	json := res.JSON200
	data := [][]string{
		locationutils.ConvertLocationJSONToTableRow(json),
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
