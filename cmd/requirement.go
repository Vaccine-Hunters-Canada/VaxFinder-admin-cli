package cmd

import (
	"github.com/fatih/color"
	"vf-admin/internal/api/requirement/add"
	"vf-admin/internal/api/requirement/get"
	"vf-admin/internal/api/requirement/list"
	"vf-admin/internal/api/requirement/remove"
	"vf-admin/internal/api/requirement/update"

	"vf-admin/internal/cmdrun"
	"vf-admin/internal/utils"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		var op get.HTTPOperation
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin requirement list`
var requirementListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of requirements",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var op list.HTTPOperation
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)
	},
}

// Command: `vf-admin requirement add`
var requirementAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new requirement",
	Example: heredoc.Doc(`
		# HTTPOperation a new requirement with name "18+" and description "Any individual older than 18 years of age.".
		$ vf-admin requirement add --name "18+" --description "Any individual older than 18 years of age."
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
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		var op add.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(name, description); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		// Retrieve the flags to be placed inside the HTTP body
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		var op update.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(name, description); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
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

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

func init() {
	requirementCmd.PersistentFlags().Bool("dry-run", false, "print the HTTP request that would be sent to the server as a cURL command")
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
