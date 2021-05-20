package main

import (
	"vf-admin/cmd"
	"vf-admin/internal/utils"
)

var (
	version string
	commit  string
	date    string
	baseURL string
)

func main() {
	// Set build version
	if len(version) > 0 {
		utils.SetVersion(version)
	}

	// Set API base URL
	if len(baseURL) > 0 {
		utils.SetBaseURL(baseURL)
	}

	cmd.Execute()
}
