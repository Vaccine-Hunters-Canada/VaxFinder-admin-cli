package add

import (
	"context"
	"time"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va timeslot add <va-id>`
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
	timeString, _ := flags.GetString("time")
	time, tErr := time.Parse(time.RFC3339, timeString)
	if tErr != nil {
		color.Red(tErr.Error())
		return nil
	}
	body := api.CreateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsPostJSONRequestBody{Time: time}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("add", "added", "va timeslot")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}

	_ = spinner.Start()

	res, rErr := client.CreateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsPostWithResponse(context.Background(), id, body)

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

	colNames := []string{"id", "vaccine availability", "time", "taken at", "created at"}
	json := res.JSON200
	var takenAt string
	if json.TakenAt == nil {
		takenAt = ""
	} else {
		takenAt = json.TakenAt.String()
	}
	data := [][]string{
		{
			json.Id, json.VaccineAvailability, json.Time.String(), takenAt, json.CreatedAt.String(),
		},
	}
	utils.RenderDefaultTable(colNames, data)

	return nil
}
