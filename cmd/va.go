package cmd

import (
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
	Args:  cobra.ExactArgs(1),
}

// Command: `vf-admin va list`
var vaListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of vaccine availabilities within the vicinity of a postal code",
	Args:  cobra.ExactArgs(0),
}

// Command: `vf-admin va add`
var vaAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new vaccine availability",
	Args:  cobra.ExactArgs(0),
}

// Command: `vf-admin va update <id>`
var vaUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a vaccine availability with a specified id",
	Args:  cobra.ExactArgs(1),
}

// Command: `vf-admin va remove <id>`
var vaRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a vaccine availability with a specified id",
	Args:  cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(vaCmd)

	// Get command
	vaCmd.AddCommand(vaGetCmd)

	// List command
	vaCmd.AddCommand(vaListCmd)

	// Add command
	vaCmd.AddCommand(vaAddCmd)

	// Update command
	vaCmd.AddCommand(vaUpdateCmd)

	// Remove command
	vaCmd.AddCommand(vaRemoveCmd)
}
