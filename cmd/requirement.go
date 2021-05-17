package cmd

import (
	"vf-admin/internal/cmdrun/requirement/add"
	"vf-admin/internal/cmdrun/requirement/get"
	"vf-admin/internal/cmdrun/requirement/list"

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

// Command: `vf-admin requirement add <id>`
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
	Args:  cobra.ExactArgs(1),
}

// Command: `vf-admin requirement remove <id>`
var requirementRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a requirement with a specified id",
	Args:  cobra.ExactArgs(1),
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
	requirementCmd.AddCommand(requirementUpdateCmd)

	// Remove command
	requirementCmd.AddCommand(requirementRemoveCmd)
}
