package cmd

import (
	"vf-admin/internal/cmdrun/va/requirement/add"
	"vf-admin/internal/cmdrun/va/requirement/list"
	"vf-admin/internal/cmdrun/va/requirement/remove"
	"vf-admin/internal/cmdrun/va/requirement/update"

	"github.com/MakeNowJust/heredoc"
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
	RunE: list.CmdRunE,
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
	RunE: add.CmdRunE,
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
	RunE: update.CmdRunE,
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
	RunE: remove.CmdRunE,
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
