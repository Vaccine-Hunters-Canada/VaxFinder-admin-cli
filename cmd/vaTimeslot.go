package cmd

import (
	"time"
	"vf-admin/internal/api/va/timeslot/add"
	"vf-admin/internal/api/va/timeslot/list"
	"vf-admin/internal/api/va/timeslot/remove"
	"vf-admin/internal/api/va/timeslot/update"
	"vf-admin/internal/cmdrun"
	"vf-admin/internal/utils"

	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Command: `vf-admin va timeslot`
var vaTimeslotCmd = &cobra.Command{
	Use:     "timeslot",
	Aliases: []string{"ts"},
	Short:   "Manage timeslots for vaccine availability",
}

// Command: `vf-admin va timeslot list <va-id>`
var vaTimeslotListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Retrieve a list of timeslots for a vaccine availability",
	Example: heredoc.Doc(`
			# List the timeslots for a vaccine availability.
			$ vf-admin va timeslot list <va-id>
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op list.HTTPOperation

		op.SetRequestURLArguments(args)

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin va add <va-id>`
var vaTimeslotAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new timeslot for vaccine availability",

	Example: heredoc.Doc(`
			# Add a new timeslot 10 for a vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 with time "2006-01-02T15:04:05Z"
			$ vf-admin va timeslot add c7bc794c-9905-4588-81e6-e557e1a494c4 --time "2006-01-02T15:04:05Z"
	`),

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op add.HTTPOperation

		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}
		op.SetAuthKey(key)

		// Retrieve the flags to be placed inside the HTTP body
		flags := cmd.Flags()

		timeString, _ := flags.GetString("time")
		time, tErr := time.Parse(time.RFC3339, timeString)
		if tErr != nil {
			color.Red(tErr.Error())
			return nil
		}

		if err := op.SetRequestBody(time); err != nil {
			return err
		}

		op.SetRequestURLArguments(args)

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin va update <va-id> <timeslot-id>`
var vaTimeslotUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up"},
	Short:   "Update a timeslot for vaccine availability with specified ids",
	/*
		Example: heredoc.Doc(`
				# update vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and timeslot id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 with time "2006-01-02T15:04:05Z07:00"
				$ vf-admin va timeslot update c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 --time "2006-01-02T15:04:05Z07:00"
		`),
	*/
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op update.HTTPOperation

		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		if err := op.SetRequestURLArguments(args); err != nil {
			return err
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

		op.SetAuthKey(key)

		if err := op.SetRequestBody(slotTime, &takenAtBody); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin va remove <va-id> <timeslot-id>`
var vaTimeslotRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a timeslot from a vaccine availability with specified ids",
	Example: heredoc.Doc(`
			# Remove vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and timeslot id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82
			$ vf-admin va timeslot remove c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82
	`),
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op remove.HTTPOperation

		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		op.SetRequestURLArguments(args)
		op.SetAuthKey(key)

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

func init() {
	vaCmd.AddCommand(vaTimeslotCmd)

	// List command
	vaTimeslotCmd.AddCommand(vaTimeslotListCmd)

	// Add command
	vaTimeslotAddCmd.Flags().String("time", "", "time of the slot (RFC 3339)")
	_ = vaTimeslotAddCmd.MarkFlagRequired("time")
	vaTimeslotCmd.AddCommand(vaTimeslotAddCmd)

	// Update command
	vaTimeslotUpdateCmd.Flags().String("time", "", "time of the slot (RFC 3339)")
	_ = vaTimeslotUpdateCmd.MarkFlagRequired("time")
	vaTimeslotUpdateCmd.Flags().String("takenAt", "", "time of the slot was taken (RFC 3339)")
	vaTimeslotCmd.AddCommand(vaTimeslotUpdateCmd)

	// Remove command
	vaTimeslotCmd.AddCommand(vaTimeslotRemoveCmd)
}
