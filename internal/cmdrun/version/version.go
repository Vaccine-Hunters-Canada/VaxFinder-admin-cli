package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"vf-admin/internal/utils"
)

var repo = "https://github.com/Vaccine-Hunters-Canada/VaxFinder-admin-cli"

// CmdRun is what's executed when running `vf-admin version`
func CmdRun(cmd *cobra.Command, args []string) {
	version, tag, date := utils.GetBuildInfo()

	if tag == "unknown" && date == "unknown" {
		fmt.Println("vf-admin/" + version)
		fmt.Println(repo)
	} else {
		fmt.Println("vf-admin/v" + version + " (" + date + ")")
		fmt.Println("https://github.com/Vaccine-Hunters-Canada/VaxFinder-admin-cli/releases/tag/" + tag)
	}
}
