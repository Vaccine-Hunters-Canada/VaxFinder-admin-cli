package main

import (
	"github.com/spf13/cobra/doc"
	"log"
	"os"
	"path/filepath"
	"vf-admin/cmd"
)

var buildVersion = "source"

func main() {
	if len(os.Args) > 1 {
		// Remove `v` prefix if one exists
		if len(os.Args[1]) > 1 && os.Args[1][0] == 'v' {
			buildVersion = os.Args[1][1:]
		} else {
			buildVersion = os.Args[1]
		}
	}

	// Create the necessary directories based on version
	folderPath := filepath.Join("./docs", buildVersion)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Generate markdown docs in `folderPath`
	cmd.RootCmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTree(cmd.RootCmd, folderPath); err != nil {
		log.Fatal(err)
	}
}
