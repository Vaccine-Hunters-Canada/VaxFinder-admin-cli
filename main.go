package main

import (
	"vf-admin/cmd"
	"vf-admin/internal/utils"
)

var buildVersion string
var baseURL string

func main() {
	// Set build version
	if len(buildVersion) > 0 {
		utils.SetVersion(buildVersion)
	}

	// Set API base URL
	if len(baseURL) > 0 {
		utils.SetBaseURL(baseURL)
	}

	cmd.Execute()
}
