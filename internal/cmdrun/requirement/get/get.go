package get

import (
	"vf-admin/internal/api/requirement"
	"vf-admin/internal/cmdrun"

	"github.com/spf13/cobra"
)

// CmdRunE is what's executed when running `vf-admin requirement get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	var op requirement.Get
	if err := op.SetRequestURLArguments(args); err != nil {
		return err
	}
	cmdrun.HTTPOpCmdRun(op)

	return nil
}
