package list

import (
	"context"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va timelist list <va-id>>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	id := args[0]

	// Create the API client
	client, cErr := utils.GetAPIClient()
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("list", "got", "va requirement")
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.ListTimeslotsForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsGetWithResponse(context.Background(), id)

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
	var data [][]string
	for _, row := range *res.JSON200 {
		var takenAt string
		if row.TakenAt == nil {
			takenAt = ""
		} else {
			takenAt = row.TakenAt.String()
		}
		data = append(data, []string{
			row.Id, row.VaccineAvailability, row.Time.String(), takenAt, row.CreatedAt.String(),
		})
	}

	utils.RenderDefaultTable(colNames, data)

	return nil
}
