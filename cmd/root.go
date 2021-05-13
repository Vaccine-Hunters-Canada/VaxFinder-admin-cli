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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Should be overwritten
var version = "0.0.0"

// rootCmd represents the base command when called without any subcommands
// Command: vf-admin
var rootCmd = &cobra.Command{
	Use:   "vf-admin",
	Short: "vf-admin is a CLI to manage vaccine availabilities and other data for the Vaccine Hunters Finder tool.",
	Long: ``,
	Version: version,
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().String("profile", "default", "specifies the named profile to use for this command")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// initConfig reads in config file.
func initConfig() {
	// Configuration file in home directory with name ".vf-admin.json"

	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Configuration file path
	const configName = ".vf-admin.json"
	configPath := path.Join(home, configName)
	viper.SetConfigType("json")

	// If the config file doesn't exist, then create it.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = viper.WriteConfigAs(configPath)
		cobra.CheckErr(err)
	}

	// Search config in home directory with name ".vf-admin.json".
	viper.SetConfigFile(configPath)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
