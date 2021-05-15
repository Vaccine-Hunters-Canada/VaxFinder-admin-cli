package utils

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetKeyFromProfile retrieves an authentication key stored in the configuration file under a specified profile
func GetKeyFromProfile(cmd *cobra.Command) (string, error) {
	profile, _ := cmd.Flags().GetString("profile")

	// If profile doesn't exist
	if !viper.IsSet(profile) {
		return "", errors.New("profile doesn't exist")
	}

	// If key in profile doesn't exist
	if !viper.IsSet(profile + ".key") {
		return "", errors.New("key isn't set")
	}

	return viper.GetString(profile + ".key"), nil
}
