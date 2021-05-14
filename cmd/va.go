package cmd

import (
	"errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Command: vf-admin va
var vaCmd = &cobra.Command{
	Use:   "va",
	Short: "Manage vaccine availabilities",
	Long:  ``,
}

// Command: vf-admin va get
var vaRetrieveCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a vaccine availability with a specified id",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an id as an argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va list
var vaListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of vaccine availabilities within the vicinity of a postal code",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a postal code argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va add
var vaAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new vaccine availability",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va update
var vaUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a vaccine availability with a specified id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va remove
var vaRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a vaccine availability with a specified id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

func init() {
	rootCmd.AddCommand(vaCmd)
	vaCmd.AddCommand(vaRetrieveCmd)
	vaCmd.AddCommand(vaListCmd)
	vaCmd.AddCommand(vaAddCmd)
	vaCmd.AddCommand(vaUpdateCmd)
	vaCmd.AddCommand(vaRemoveCmd)
}
