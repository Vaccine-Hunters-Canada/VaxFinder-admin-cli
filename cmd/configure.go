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
	RunE:  configureCmdRun,
}

func init() {
	RootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringP("key", "k", "", "The authentication key for future requests.")
	_ = configureCmd.MarkFlagRequired("key")
}

func configureCmdRun(cmd *cobra.Command, args []string) error {
	profile, _ := cmd.Flags().GetString("profile")
	key, _ := cmd.Flags().GetString("key")

	// Return error if the key is not a UUID
	if _, keyErr := uuid.Parse(key); keyErr != nil {
		return errors.New("the key provided is not valid")
	}

	// TODO: Make REST call to verify key and get role type

	// Update the configuration file for the specific profile
	viper.Set(profile+".key", key)
	viper.Set(profile+".role", 0)
	viper.Set(profile+".updated", time.Now())
	if writeErr := viper.WriteConfig(); writeErr != nil {
		return writeErr
	}

	color.Green("Successfully updated configuration for the `%s` profile.\n", profile)
	return nil
}
