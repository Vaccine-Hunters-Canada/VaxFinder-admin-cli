package cmd

import (
	"vf-admin/internal/cmdrun/organization/add"
	"vf-admin/internal/cmdrun/organization/get"
	"vf-admin/internal/cmdrun/organization/list"
	"vf-admin/internal/cmdrun/organization/remove"
	"vf-admin/internal/cmdrun/organization/update"

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
	Short: "Retrieve an organization with a specified id",
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
	Example: heredoc.Doc(`
			# Add a new organization with short name "WHO", full name "World Health Organization", description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." and url "https://www.who.int/"
			$ vf-admin organization add --shortName WHO --fullName "World Health Organization" --description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." --url "https://www.who.int/"
	`),
	Args: cobra.ExactArgs(0),
	RunE: add.CmdRunE,
}

// Command: `vf-admin organization update <id>`
var organizationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an organization with a specified id",
	Example: heredoc.Doc(`
			# Update the organization with id 20 to have short name "WHO", full name "World Health Organization", description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." and url "https://www.who.int/"
			$ vf-admin organization update 20 --shortName WHO --fullName "World Health Organization" --description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." --url "https://www.who.int/"
	`),
	Args: cobra.ExactArgs(1),
	RunE: update.CmdRunE,
}

// Command: `vf-admin organization remove <id>`
var organizationRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an organization with a specified id",
	Example: heredoc.Doc(`
			# Remove the organization with id 20.
			$ vf-admin organization remove 20
	`),
	Args: cobra.ExactArgs(1),
	RunE: remove.CmdRunE,
}

func init() {
	RootCmd.AddCommand(organizationCmd)

	// Get command
	organizationCmd.AddCommand(organizationGetCmd)

	// List command
	organizationCmd.AddCommand(organizationListCmd)

	// Add command
	organizationAddCmd.Flags().String("shortName", "", "short name of the organization")
	_ = organizationAddCmd.MarkFlagRequired("shortName")
	organizationAddCmd.Flags().String("fullName", "", "full name of the organization")
	organizationAddCmd.Flags().String("description", "", "description of organization")
	organizationAddCmd.Flags().String("url", "", "url of organization")
	organizationCmd.AddCommand(organizationAddCmd)

	// Update command
	organizationUpdateCmd.Flags().String("shortName", "", "short name of the organization")
	_ = organizationUpdateCmd.MarkFlagRequired("shortName")
	organizationUpdateCmd.Flags().String("fullName", "", "full name of the organization")
	organizationUpdateCmd.Flags().String("description", "", "description of organization")
	organizationUpdateCmd.Flags().String("url", "", "url of organization")
	organizationCmd.AddCommand(organizationUpdateCmd)

	// Remove command
	organizationCmd.AddCommand(organizationRemoveCmd)
}
