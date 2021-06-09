package cmd

import (
	"github.com/fatih/color"
	"vf-admin/internal/api/organization/add"
	"vf-admin/internal/api/organization/get"
	"vf-admin/internal/api/organization/list"
	"vf-admin/internal/api/organization/remove"
	"vf-admin/internal/api/organization/update"
	"vf-admin/internal/cmdrun"
	"vf-admin/internal/utils"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		var op get.HTTPOperation
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		cmdrun.RunHTTPOperation(op)

		return nil
	},
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
	Run: func(cmd *cobra.Command, args []string) {
		var op list.HTTPOperation
		cmdrun.RunHTTPOperation(op)
	},
}

// Command: `vf-admin organization add`
var organizationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new organization",
	Example: heredoc.Doc(`
			# Add a new organization with short name "WHO", full name "World Health Organization", description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." and url "https://www.who.int/"
			$ vf-admin organization add --shortName WHO --fullName "World Health Organization" --description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." --url "https://www.who.int/"
	`),
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		// Retrieve the flags to be placed inside the HTTP body
		flags := cmd.Flags()
		shortName, _ := flags.GetString("shortName")
		var fullName, description, url *string
		if flags.Changed("fullName") {
			t, _ := flags.GetString("fullName")
			fullName = &t
		} else {
			fullName = nil
		}
		if flags.Changed("description") {
			t, _ := flags.GetString("description")
			description = &t
		} else {
			description = nil
		}
		if flags.Changed("url") {
			t, _ := flags.GetString("url")
			url = &t
		} else {
			url = nil
		}

		var op add.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(shortName, fullName, description, url); err != nil {
			return err
		}
		cmdrun.RunHTTPOperation(op)

		return nil
	},
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		// Retrieve the flags to be placed inside the HTTP body
		flags := cmd.Flags()
		shortName, _ := flags.GetString("shortName")
		var fullName, description, url *string
		if flags.Changed("fullName") {
			t, _ := flags.GetString("fullName")
			fullName = &t
		} else {
			fullName = nil
		}
		if flags.Changed("description") {
			t, _ := flags.GetString("description")
			description = &t
		} else {
			description = nil
		}
		if flags.Changed("url") {
			t, _ := flags.GetString("url")
			url = &t
		} else {
			url = nil
		}

		var op update.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(shortName, fullName, description, url); err != nil {
			return err
		}
		cmdrun.RunHTTPOperation(op)

		return nil
	},
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		var op remove.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		cmdrun.RunHTTPOperation(op)

		return nil
	},
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
