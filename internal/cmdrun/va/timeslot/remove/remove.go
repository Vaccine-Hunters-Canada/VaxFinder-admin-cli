package remove

import (
	"context"
	"fmt"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin va timeslot remove <va-id> <timeslot-id>`
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
	client, cErr := api.GetAPIClientFromKey(key)
	if cErr != nil {
		color.Red(cErr.Error())
		return nil
	}

	// Create spinner
	spinner, sErr := utils.GetDefaultSpinnerForHTTPOp("remove", "removed", fmt.Sprint("vaccine availability ", vaID, " timeslot id ", slotID))
	if sErr != nil {
		color.Red(sErr.Error())
		return nil
	}
	_ = spinner.Start()

	res, rErr := client.DeleteTimeslotForVaccineAvailabilityByIdApiV1VaccineAvailabilityVaccineAvailabilityIdTimeslotsTimeslotIdDeleteWithResponse(context.Background(), vaID, slotID)

	if rErr != nil {
		spinner.StopFailMessage(rErr.Error())
		_ = spinner.StopFail()
		return nil
	}

	if res.StatusCode() != 204 {
		spinner.StopFailMessage(res.Status() + ": " + string(res.Body))
		_ = spinner.StopFail()
		return nil
	}

	_ = spinner.Stop()

	return nil
}
