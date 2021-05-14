package main

import (
	"vf-admin/cmd"
	"vf-admin/internal/utils"
)

var buildVersion string

func main() {
	// Set build version
	if len(buildVersion) > 0 {
		utils.SetVersion(buildVersion)
	}

	cmd.Execute()
}
