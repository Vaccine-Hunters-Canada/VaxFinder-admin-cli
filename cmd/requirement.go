package cmd

import (
	"vf-admin/internal/cmdrun/requirement/add"
	"vf-admin/internal/cmdrun/requirement/get"
	"vf-admin/internal/cmdrun/requirement/list"
	"vf-admin/internal/cmdrun/requirement/remove"
	"vf-admin/internal/cmdrun/requirement/update"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin requirement`
var requirementCmd = &cobra.Command{
	Use:   "requirement",
	Short: "Manage requirements",
}

// Command: `vf-admin requirement get <id>`
var requirementGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a requirement with a specified id",
	Example: heredoc.Doc(`
			# Get the requirement with id 1.
			$ vf-admin requirement get 1
	`),
	Args: cobra.ExactArgs(1),
	RunE: get.CmdRunE,
}

// Command: `vf-admin requirement list`
var requirementListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of requirements",
	Args:  cobra.ExactArgs(0),
	Run:   list.CmdRun,
}

// Command: `vf-admin requirement add`
var requirementAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new requirement",
	Example: heredoc.Doc(`
			# Add a new requirement with name "18+" and description "Any individual older than 18 years of age.".
			$ vf-admin requirement add --name "18+" --description "Any individual older than 18 years of age."
	`),
	Args: cobra.ExactArgs(0),
	RunE: add.CmdRunE,
}

// Command: `vf-admin requirement update <id>`
var requirementUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a requirement with a specified id",
	Example: heredoc.Doc(`
			# Update the requirement with id 8 to have name "High-Risk" and description "Highest- and High-Risk Health Conditions.".
			$ vf-admin requirement update 8 --name "High-Risk" --description "Highest- and High-Risk Health Conditions."
	`),
	Args: cobra.ExactArgs(1),
	RunE: update.CmdRunE,
}

// Command: `vf-admin requirement remove <id>`
var requirementRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a requirement with a specified id",
	Example: heredoc.Doc(`
			# Remove the requirement with id 8.
			$ vf-admin requirement remove 8
	`),
	Args: cobra.ExactArgs(1),
	RunE: remove.CmdRunE,
}

func init() {
	RootCmd.AddCommand(requirementCmd)

	// Get command
	requirementCmd.AddCommand(requirementGetCmd)

	// List command
	requirementCmd.AddCommand(requirementListCmd)

	// Add command
	requirementAddCmd.Flags().String("name", "", "name of requirement")
	_ = requirementAddCmd.MarkFlagRequired("name")
	requirementAddCmd.Flags().String("description", "", "description of requirement")
	_ = requirementAddCmd.MarkFlagRequired("description")
	requirementCmd.AddCommand(requirementAddCmd)

	// Update command
	requirementUpdateCmd.Flags().String("name", "", "name of requirement")
	_ = requirementUpdateCmd.MarkFlagRequired("name")
	requirementUpdateCmd.Flags().String("description", "", "description of requirement")
	_ = requirementUpdateCmd.MarkFlagRequired("description")
	requirementCmd.AddCommand(requirementUpdateCmd)

	// Remove command
	requirementCmd.AddCommand(requirementRemoveCmd)
}
