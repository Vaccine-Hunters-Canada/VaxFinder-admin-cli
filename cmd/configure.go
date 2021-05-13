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
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// Command: vf-admin configure
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure a named profile with settings",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, _ := cmd.Flags().GetString("profile")
		key, _ := cmd.Flags().GetString("key")

		// Return error if the key is not a UUID
		if _, keyErr := uuid.Parse(key); keyErr != nil {
			return errors.New("the key provided is not valid")
		}

		// TODO: Make REST call to verify key and get role type

		// Update the configuration file for the specific profile
		viper.Set(profile + ".key", key)
		viper.Set(profile + ".role", 0)
		viper.Set(profile + ".updated", time.Now())
		if writeErr := viper.WriteConfig(); writeErr != nil {
			return writeErr
		}

		color.Green("Successfully updated configuration for the `%s` profile.\n", profile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command is called directly
	configureCmd.Flags().StringP("key", "k", "", "The authentication key for future requests.")
	configureCmd.MarkFlagRequired("key")
}
