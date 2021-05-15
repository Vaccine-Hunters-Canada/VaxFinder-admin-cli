package main

import (
	"github.com/spf13/cobra/doc"
	"log"
	"vf-admin/cmd"
)

func main() {
	cmd.RootCmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTree(cmd.RootCmd, "./docs"); err != nil {
		log.Fatal(err)
	}
}
