package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"vf-admin/internal/cmdrun/configure"
)

// Command: `vf-admin configure`
var configureCmd = &cobra.Command{
	Use:   "configure",
	Args:  cobra.ExactArgs(0),
	Short: "Configure a named profile with credentials",
	Long: heredoc.Doc(`
			Configure the default profile or simply create new named profiles with an authentication key.

			A named profile is a set of settings and credentials that you can apply to a vf-admin CLI command.
	`),
	Example: heredoc.Doc(`z
			# set up authentication key for default profile
			$ vf-admin configure --key 7260841f-b47a-4b5c-9830-585af07c4405

			# set up authentication key for a custom profile
			$ vf-admin configure --profile alvin --key 7260841f-b47a-4b5c-9830-585af07c4405
	`),
	RunE: configure.CmdRunE,
}

func init() {
	RootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringP("key", "k", "", "The authentication key for future requests.")
	_ = configureCmd.MarkFlagRequired("key")
}
