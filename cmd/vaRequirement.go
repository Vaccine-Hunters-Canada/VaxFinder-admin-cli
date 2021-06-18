package cmd

import (
	"vf-admin/internal/api/va/requirement/add"
	"vf-admin/internal/api/va/requirement/list"
	"vf-admin/internal/api/va/requirement/remove"
	"vf-admin/internal/api/va/requirement/update"
	"vf-admin/internal/cmdrun"
	"vf-admin/internal/utils"

	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Command: `vf-admin va requirement`
var vaRequirementCmd = &cobra.Command{
	Use:   "requirement",
	Short: "Manage requirements for vaccine availability",
}

// Command: `vf-admin va requirement list <va-id>`
var vaRequirementListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of requirements for a vaccine availability",
	Example: heredoc.Doc(`
			# List the requirements for a vaccine availability.
			$ vf-admin va requirement list <va-id>
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op list.HTTPOperation

		if err:= op.SetRequestURLArguments(args); err!= nil {
			return err
		}
		cmdrun.RunHTTPOperation(op)
		return nil
	},
}

// Command: `vf-admin va add <va-id>`
var vaRequirementAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new requirement for vaccine availability",
	Example: heredoc.Doc(`
			# Add a new requirement 10 for a vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4
			$ vf-admin va requirement add c7bc794c-9905-4588-81e6-e557e1a494c4 --requirement 10
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

		requirement, _ := flags.GetInt("requirement")

		var op add.HTTPOperation
		op.SetAuthKey(key)

		if err := op.SetRequestBody(requirement); err != nil {
			return err
		}

		op.SetRequestURLArguments(args)
		cmdrun.RunHTTPOperation(op)

		return nil
	},
}

// Command: `vf-admin va update <va-id> <requirement-id>`
var vaRequirementUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a requirement for vaccine availability with specified ids",
	Example: heredoc.Doc(`
			# update vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and requirement id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 with requirement id 10 and active false
			$ vf-admin va requirement update c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 --requirement 10 --active=false
	`),
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
		requirement, _ := flags.GetInt("requirement")
		active, _ := flags.GetBool("active")

		op.SetAuthKey(key)

		if err := op.SetRequestBody(requirement, active); err != nil {
			return err
		}

		cmdrun.RunHTTPOperation(op)

		return nil
	},
}

// Command: `vf-admin va remove <va-id> <requirement-id>`
var vaRequirementRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a requirement from a vaccine availability with specified ids",
	Example: heredoc.Doc(`
			# Remove vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and requirement id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82
			$ vf-admin va requirement remove c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82
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

		cmdrun.RunHTTPOperation(op)

		return nil
	},
}

func init() {
	vaCmd.AddCommand(vaRequirementCmd)

	// List command
	vaRequirementCmd.AddCommand(vaRequirementListCmd)

	// Add command
	vaRequirementAddCmd.Flags().Int("requirement", 0, "id of requirement")
	_ = vaRequirementAddCmd.MarkFlagRequired("requirement")
	vaRequirementCmd.AddCommand(vaRequirementAddCmd)

	// Update command
	vaRequirementUpdateCmd.Flags().Int("requirement", 0, "id of requirement")
	_ = vaRequirementUpdateCmd.MarkFlagRequired("requirement")
	vaRequirementUpdateCmd.Flags().Bool("active", false, "if the requirement is active or not")
	_ = vaRequirementUpdateCmd.MarkFlagRequired("active")
	vaRequirementCmd.AddCommand(vaRequirementUpdateCmd)

	// Remove command
	vaRequirementCmd.AddCommand(vaRequirementRemoveCmd)
}
