package add

import (
	"context"
	"strconv"
	"time"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va add`
func CmdRunE(cmd *cobra.Command, args []string) error {
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
	t, _ := flags.GetString("date")
	date, tErr := time.Parse("2006-01-02", t)
	if tErr != nil {
		color.Red(tErr.Error())
		return nil
	}

	location, _ := flags.GetInt("location")
	inputTypeInt, _ := flags.GetInt("inputtype")
	inputType := api.InputTypeEnum(inputTypeInt)
	numberAvailable, _ := flags.GetInt("numberavailable")

	var numberTotal, vaccine *int
	var tags *string

	if flags.Changed("numberTotal") {
		t, _ := flags.GetInt("numberTotal")
		numberTotal = &t
	} else {
		numberTotal = nil
	}
	if flags.Changed("vaccine") {
		t, _ := flags.GetInt("vaccine")
		vaccine = &t
	} else {
		vaccine = nil
	}
	if flags.Changed("tags") {
		t, _ := flags.GetString("tags")
		tags = &t
	} else {
		tags = nil
	}

	body := api.CreateVaccineAvailabilityApiV1VaccineAvailabilityPostJSONRequestBody{Date: date, Location: location, InputType: inputType, NumberAvailable: numberAvailable, NumberTotal: numberTotal, Vaccine: vaccine, Tags: tags}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("add", "added", "va")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.CreateVaccineAvailabilityApiV1VaccineAvailabilityPostWithResponse(context.Background(), body)

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

	colNames := []string{"id", "date", "number available", "number total", "vaccine", "input type", "tags", "location", "created at"}
	json := res.JSON200
	data := [][]string{
		{
			json.Id, json.Date.String(), strconv.Itoa(json.NumberAvailable), utils.CoalesceInt(json.NumberTotal), utils.CoalesceInt(json.Vaccine), strconv.Itoa(int(json.InputType)), utils.CoalesceString(json.Tags), strconv.Itoa(json.Location), json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
