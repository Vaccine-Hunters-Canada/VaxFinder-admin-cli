package cmd

import (
	"github.com/fatih/color"
	"vf-admin/internal/api/address/add"
	"vf-admin/internal/api/address/get"
	"vf-admin/internal/api/address/list"
	"vf-admin/internal/api/address/remove"
	"vf-admin/internal/api/address/update"
	"vf-admin/internal/cmdrun"
	"vf-admin/internal/utils"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin address`
var addressCmd = &cobra.Command{
	Use:     "address",
	Aliases: []string{"addr"},
	Short:   "Manage addresses",
}

// Command: `vf-admin address get <id>`
var addressGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an address with a specified id",
	Example: heredoc.Doc(`
			# Get the address with id 1.
			$ vf-admin address get 1
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var op get.HTTPOperation
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin address list`
var addressListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Retrieve a list of addresses",
	Example: heredoc.Doc(`
			# List the addresses.
			$ vf-admin address list
	`),
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var op list.HTTPOperation
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)
	},
}

// Command: `vf-admin address add`
var addressAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new address",
	Example: heredoc.Doc(`
			# Add a new address with province "Ontario", postal code "K1A0A9", latitude "45.424807" and longitude "-75.699234"
			$ vf-admin address add --province "Ontario" --postcode "K1A0A9" --latitude "45.424807" --longitude "-75.699234"
	`),
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		// Retrieve the flags to be placed inside the HTTP body
		flags := cmd.Flags()
		province, _ := flags.GetString("province")
		postcode, _ := flags.GetString("postcode")
		latitude, _ := flags.GetFloat32("latitude")
		longitude, _ := flags.GetFloat32("longitude")
		var line1, line2, city *string
		if flags.Changed("line1") {
			t, _ := flags.GetString("line1")
			line1 = &t
		} else {
			line1 = nil
		}
		if flags.Changed("line2") {
			t, _ := flags.GetString("line2")
			line2 = &t
		} else {
			line2 = nil
		}
		if flags.Changed("city") {
			t, _ := flags.GetString("city")
			city = &t
		} else {
			city = nil
		}

		var op add.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(province, postcode, latitude, longitude, line1, line2, city); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin address update <id>`
var addressUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up"},
	Short:   "Update an address with a specified id",
	Example: heredoc.Doc(`
			# Update the address with id 20 to have province "Ontario", postal code "K1A0A9", latitude "45.424807" and longitude "-75.699234"
			$ vf-admin address update 20 --province "Ontario" --postcode "K1A0A9" --latitude "45.424807" --longitude "-75.699234"
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		// Retrieve the flags to be placed inside the HTTP body
		flags := cmd.Flags()
		province, _ := flags.GetString("province")
		postcode, _ := flags.GetString("postcode")
		latitude, _ := flags.GetFloat32("latitude")
		longitude, _ := flags.GetFloat32("longitude")
		var line1, line2, city *string
		if flags.Changed("line1") {
			t, _ := flags.GetString("line1")
			line1 = &t
		} else {
			line1 = nil
		}
		if flags.Changed("line2") {
			t, _ := flags.GetString("line2")
			line2 = &t
		} else {
			line2 = nil
		}
		if flags.Changed("city") {
			t, _ := flags.GetString("city")
			city = &t
		} else {
			city = nil
		}

		var op update.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}
		if err := op.SetRequestBody(province, postcode, latitude, longitude, line1, line2, city); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

// Command: `vf-admin address remove <id>`
var addressRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove an address with a specified id",
	Example: heredoc.Doc(`
			# Remove the address with id 20.
			$ vf-admin address remove 20
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the authentication key from configuration file
		key, kErr := utils.GetKeyFromProfile(cmd)
		if kErr != nil {
			color.Red(kErr.Error())
			return nil
		}

		var op remove.HTTPOperation
		op.SetAuthKey(key)
		if err := op.SetRequestURLArguments(args); err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		cmdrun.RunHTTPOperation(op, dryRun)

		return nil
	},
}

func init() {
	addressCmd.PersistentFlags().Bool("dry-run", false, "print the HTTP request that would be sent to the server as a cURL command")
	RootCmd.AddCommand(addressCmd)

	// Get command
	addressCmd.AddCommand(addressGetCmd)

	// List command
	addressCmd.AddCommand(addressListCmd)

	// Add command
	addressAddCmd.Flags().String("line1", "", "line 1 of the address")
	addressAddCmd.Flags().String("line2", "", "line 2 of the address")
	addressAddCmd.Flags().String("city", "", "city of the address")
	addressAddCmd.Flags().String("province", "", "province of the address")
	_ = addressAddCmd.MarkFlagRequired("province")
	addressAddCmd.Flags().String("postcode", "", "postal code of the address")
	_ = addressAddCmd.MarkFlagRequired("postcode")
	addressAddCmd.Flags().Float32("latitude", 0, "latitude code of the address")
	_ = addressAddCmd.MarkFlagRequired("latitude")
	addressAddCmd.Flags().Float32("longitude", 0, "longitude code of the address")
	_ = addressAddCmd.MarkFlagRequired("longitude")
	addressCmd.AddCommand(addressAddCmd)

	// Update command
	addressUpdateCmd.Flags().String("line1", "", "line 1 of the address")
	addressUpdateCmd.Flags().String("line2", "", "line 2 of the address")
	addressUpdateCmd.Flags().String("city", "", "city of the address")
	addressUpdateCmd.Flags().String("province", "", "province of the address")
	_ = addressUpdateCmd.MarkFlagRequired("province")
	addressUpdateCmd.Flags().String("postcode", "", "postal code of the address")
	_ = addressUpdateCmd.MarkFlagRequired("postcode")
	addressUpdateCmd.Flags().Float32("latitude", 0, "latitude code of the address")
	_ = addressUpdateCmd.MarkFlagRequired("latitude")
	addressUpdateCmd.Flags().Float32("longitude", 0, "longitude code of the address")
	_ = addressUpdateCmd.MarkFlagRequired("longitude")
	addressCmd.AddCommand(addressUpdateCmd)

	// Remove command
	addressCmd.AddCommand(addressRemoveCmd)
}
