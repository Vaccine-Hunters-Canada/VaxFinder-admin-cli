package cmd

import (
	"vf-admin/internal/cmdrun/va/add"
	"vf-admin/internal/cmdrun/va/get"
	"vf-admin/internal/cmdrun/va/list"
	"vf-admin/internal/cmdrun/va/remove"
	"vf-admin/internal/cmdrun/va/update"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin va`
var vaCmd = &cobra.Command{
	Use:   "va",
	Short: "Manage vaccine availabilities",
}

// Command: `vf-admin va get <id>`
var vaGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a vaccine availability with a specified id",
	Example: heredoc.Doc(`
			# Get the vaccine availability with id 014cc133-484f-4320-be3b-444e758b64a7
			$ vf-admin va get 014cc133-484f-4320-be3b-444e758b64a7
	`),
	Args: cobra.ExactArgs(1),
	RunE: get.CmdRunE,
}

// Command: `vf-admin va list`
var vaListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of vaccine availabilities within the vicinity of a postal code",
	Example: heredoc.Doc(`
			# List the vaccine availabilities within the vicinity of a postal code.
			$ vf-admin va list --postcode K1A
	`),
	Args: cobra.ExactArgs(0),
	RunE: list.CmdRunE,
}

// Command: `vf-admin va add`
var vaAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new vaccine availability",
	Example: heredoc.Doc(`
			# Add new vaccine availability on 2021-05-25 with 3 available input type 1 location 1651 and tags vhc
			$ vf-admin va add --date "2021-05-25" --numberavailable 3 --inputtype 1 --location 1651 --tags vhc
	`),
	Args: cobra.ExactArgs(0),
	RunE: add.CmdRunE,
}

// Command: `vf-admin va update <id>`
var vaUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a vaccine availability with a specified id",
	Example: heredoc.Doc(`
			# Update vaccine availability id 7d7488e4-cc26-434d-85c4-b7df2f7e3171to be on 2021-05-25 with 3 available input type 1 location 1651 and tags vhc
			$ vf-admin va update 7d7488e4-cc26-434d-85c4-b7df2f7e3171 --date "2021-05-25" --numberavailable 3 --inputtype 1 --location 1651 --tags vhc
	`),
	Args: cobra.ExactArgs(1),
	RunE: update.CmdRunE,
}

// Command: `vf-admin va remove <id>`
var vaRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a vaccine availability with a specified id",
	Example: heredoc.Doc(`
			# Remove vaccine availability id 7d7488e4-cc26-434d-85c4-b7df2f7e3171
			$ vf-admin va remove 7d7488e4-cc26-434d-85c4-b7df2f7e3171
	`),
	Args: cobra.ExactArgs(1),
	RunE: remove.CmdRunE,
}

func init() {
	RootCmd.AddCommand(vaCmd)

	// Get command
	vaCmd.AddCommand(vaGetCmd)

	// List command
	vaListCmd.Flags().String("postcode", "", "postal code to search around")
	_ = vaListCmd.MarkFlagRequired("postcode")
	vaListCmd.Flags().String("mindate", "", "minimum date for availability (YYYY-MM-DD)")
	vaCmd.AddCommand(vaListCmd)

	// Add command
	vaAddCmd.Flags().String("date", "", "date for availability (YYYY-MM-DD)")
	_ = vaAddCmd.MarkFlagRequired("date")
	vaAddCmd.Flags().Int("location", 0, "id of the location")
	_ = vaAddCmd.MarkFlagRequired("location")
	vaAddCmd.Flags().Int("inputtype", 0, "input type")
	_ = vaAddCmd.MarkFlagRequired("inputtype")
	vaAddCmd.Flags().Int("numberavailable", 0, "number of vaccines available")
	_ = vaAddCmd.MarkFlagRequired("numberavailable")
	vaAddCmd.Flags().Int("numbertotal", 0, "total number of vaccines")
	vaAddCmd.Flags().Int("vaccine", 0, "id of the vaccine")
	vaAddCmd.Flags().String("tags", "", "tags")
	vaCmd.AddCommand(vaAddCmd)

	// Update command
	vaUpdateCmd.Flags().String("date", "", "date for availability (YYYY-MM-DD)")
	_ = vaUpdateCmd.MarkFlagRequired("date")
	vaUpdateCmd.Flags().Int("location", 0, "id of the location")
	_ = vaUpdateCmd.MarkFlagRequired("location")
	vaUpdateCmd.Flags().Int("inputtype", 0, "input type")
	_ = vaUpdateCmd.MarkFlagRequired("inputtype")
	vaUpdateCmd.Flags().Int("numberavailable", 0, "number of vaccines available")
	_ = vaUpdateCmd.MarkFlagRequired("numberavailable")
	vaUpdateCmd.Flags().Int("numbertotal", 0, "total number of vaccines")
	vaUpdateCmd.Flags().Int("vaccine", 0, "id of the vaccine")
	vaUpdateCmd.Flags().String("tags", "", "tags")
	vaCmd.AddCommand(vaUpdateCmd)

	// Remove command
	vaCmd.AddCommand(vaRemoveCmd)
}
