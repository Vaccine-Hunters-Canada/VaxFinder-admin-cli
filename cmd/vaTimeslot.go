package cmd

import (
	"vf-admin/internal/cmdrun/va/timeslot/add"
	"vf-admin/internal/cmdrun/va/timeslot/list"
	"vf-admin/internal/cmdrun/va/timeslot/remove"
	"vf-admin/internal/cmdrun/va/timeslot/update"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin va timeslot`
var vaTimeslotCmd = &cobra.Command{
	Use:   "timeslot",
	Short: "Manage timeslots for vaccine availability",
}

// Command: `vf-admin va timeslot list <va-id>`
var vaTimeslotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of timeslots for a vaccine availability",
	Example: heredoc.Doc(`
			# List the timeslots for a vaccine availability.
			$ vf-admin va timeslot list <va-id>
	`),
	Args: cobra.ExactArgs(1),
	RunE: list.CmdRunE,
}

// Command: `vf-admin va add <va-id>`
var vaTimeslotAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new timeslot for vaccine availability",
	/*
		Example: heredoc.Doc(`
				# Add a new timeslot 10 for a vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 with time "2006-01-02T15:04:05Z07:00"
				$ vf-admin va timeslot add c7bc794c-9905-4588-81e6-e557e1a494c4 --time "2006-01-02T15:04:05Z07:00"
		`),
	*/
	Args: cobra.ExactArgs(1),
	RunE: add.CmdRunE,
}

// Command: `vf-admin va update <va-id> <timeslot-id>`
var vaTimeslotUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a timeslot for vaccine availability with specified ids",
	/*
		Example: heredoc.Doc(`
				# update vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and timeslot id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 with time "2006-01-02T15:04:05Z07:00"
				$ vf-admin va timeslot update c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 --time "2006-01-02T15:04:05Z07:00"
		`),
	*/
	Args: cobra.ExactArgs(2),
	RunE: update.CmdRunE,
}

// Command: `vf-admin va remove <va-id> <timeslot-id>`
var vaTimeslotRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a timeslot from a vaccine availability with specified ids",
	Example: heredoc.Doc(`
			# Remove vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and timeslot id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 
			$ vf-admin va timeslot remove c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82
	`),
	Args: cobra.ExactArgs(2),
	RunE: remove.CmdRunE,
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
