package main

import (
	"github.com/fatih/color"
	"vf-admin/cmd"
	"vf-admin/internal/utils"
)

// ProdBaseURL is the base URL for the production API
const ProdBaseURL = "https://api.vaccinehunters.ca"

var (
	version = "source"
	tag     = "unknown"
	date    = "unknown"
	baseURL = ProdBaseURL
)

func main() {
	// Provide warning if base URL does not match ProdBaseURL
	if baseURL != ProdBaseURL {
		color.Yellow("Warning: The base URL is configured to \"" + baseURL + "\" instead of \"" + ProdBaseURL + "\".\n")
	}

	// Set build information
	utils.SetBuildInfo(version, tag, date)
	// Set API base URL
	utils.SetBaseURL(baseURL)

	cmd.Execute()
}
