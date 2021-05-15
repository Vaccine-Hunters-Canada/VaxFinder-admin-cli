package cmd

import (
	"vf-admin/internal/cmdrun/organization/get"
	"vf-admin/internal/cmdrun/organization/list"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin organization`
var organizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "Manage organizations",
}

// Command: `vf-admin organization get <id>`
var organizationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a organization with a specified id",
	Example: heredoc.Doc(`
			# Get the organization with id 1.
			$ vf-admin organization get 1
	`),
	Args: cobra.ExactArgs(1),
	RunE: get.CmdRunE,
}

// Command: `vf-admin organization list`
var organizationListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of organizations",
	Example: heredoc.Doc(`
			# List the organizations.
			$ vf-admin organization list
	`),
	Args: cobra.ExactArgs(0),
	RunE: list.CmdRunE,
}

// Command: `vf-admin organization add <id>`
var organizationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new organization",
	Args:  cobra.ExactArgs(0),
}

// Command: `vf-admin organization update <id>`
var organizationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a organization with a specified id",
	Args:  cobra.ExactArgs(1),
}

// Command: `vf-admin organization remove <id>`
var organizationRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a organization with a specified id",
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(organizationCmd)

	// Get command
	organizationCmd.AddCommand(organizationGetCmd)

	// List command
	organizationCmd.AddCommand(organizationListCmd)

	// Add command
	organizationCmd.AddCommand(organizationAddCmd)

	// Update command
	organizationCmd.AddCommand(organizationUpdateCmd)

	// Remove command
	organizationCmd.AddCommand(organizationRemoveCmd)
}
