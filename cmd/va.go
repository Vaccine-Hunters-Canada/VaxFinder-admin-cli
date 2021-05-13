/*
Copyright © 2021 VaxFinder Project

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Command: vf-admin va
var vaCmd = &cobra.Command{
	Use:   "va",
	Short: "Manage vaccine availabilities",
	Long: ``,
}

// Command: vf-admin va get
var vaRetrieveCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a vaccine availability with a specified id",
	Long: ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an id as an argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va list
var vaListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of vaccine availabilities within the vicinity of a postal code",
	Long: ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a postal code argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va add
var vaAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new vaccine availability",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va update
var vaUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a vaccine availability with a specified id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

// Command: vf-admin va remove
var vaRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a vaccine availability with a specified id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("TODO")
	},
}

func init() {
	rootCmd.AddCommand(vaCmd)
	vaCmd.AddCommand(vaRetrieveCmd)
	vaCmd.AddCommand(vaListCmd)
	vaCmd.AddCommand(vaAddCmd)
	vaCmd.AddCommand(vaUpdateCmd)
	vaCmd.AddCommand(vaRemoveCmd)
}
