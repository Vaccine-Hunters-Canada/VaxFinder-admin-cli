package cmd

import (
	"github.com/spf13/cobra"
	"vf-admin/internal/cmdrun/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of the CLI",
	Run:   version.CmdRun,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
