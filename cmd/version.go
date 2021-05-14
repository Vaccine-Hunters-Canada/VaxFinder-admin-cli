package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"vf-admin/internal/utils"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vf-admin/" + utils.GetVersion())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
