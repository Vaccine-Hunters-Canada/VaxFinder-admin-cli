package update

import (
	"context"
	"time"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va timeslot update <va-id> <timeslot-id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	vaID := args[0]
	slotID := args[1]

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
	slotTime, tErr := time.Parse(time.RFC3339, timeString)
	if tErr != nil {
		color.Red(tErr.Error())
		return nil
	}

	var takenAtBody interface{}
	if flags.Changed("takenAt") {
		t, _ := flags.GetString("takenAt")
		t2, tErr := time.Parse(time.RFC3339, t)
		if tErr != nil {
			color.Red(tErr.Error())
			return nil
		}
		takenAtBody = t2
	}
	body := api.UpdateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsTimeslotIdPutJSONRequestBody{Time: slotTime, TakenAt: &takenAtBody}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("update", "updated", "va timeslot")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.UpdateTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsTimeslotIdPutWithResponse(context.Background(), vaID, slotID, body)

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
