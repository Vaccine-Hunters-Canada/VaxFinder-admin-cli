package cmd

import (
	"vf-admin/internal/cmdrun/address/add"
	"vf-admin/internal/cmdrun/address/get"
	"vf-admin/internal/cmdrun/address/list"
	"vf-admin/internal/cmdrun/address/remove"
	"vf-admin/internal/cmdrun/address/update"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin address`
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Manage addresss",
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
	RunE: get.CmdRunE,
}

// Command: `vf-admin address list`
var addressListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of addresss",
	Example: heredoc.Doc(`
			# List the addresss.
			$ vf-admin address list
	`),
	Args: cobra.ExactArgs(0),
	RunE: list.CmdRunE,
}

// Command: `vf-admin address add`
var addressAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new address",
	Example: heredoc.Doc(`
			# Add a new address with province "Ontario", postal code "K1A0A9", latitude "45.424807" and longitude "-75.699234"
			$ vf-admin address add --province "Ontario" --postcode "K1A0A9" --latitude "45.424807"--longitude "-75.699234"
	`),
	Args: cobra.ExactArgs(0),
	RunE: add.CmdRunE,
}

// Command: `vf-admin address update <id>`
var addressUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an address with a specified id",
	Example: heredoc.Doc(`
			# Update the address with id 20 to have province "Ontario", postal code "K1A0A9", latitude "45.424807" and longitude "-75.699234"
			$ vf-admin address update 20 --province "Ontario" --postcode "K1A0A9" --latitude "45.424807"--longitude "-75.699234"
	`),
	Args: cobra.ExactArgs(1),
	RunE: update.CmdRunE,
}

// Command: `vf-admin address remove <id>`
var addressRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an address with a specified id",
	Example: heredoc.Doc(`
			# Remove the address with id 20.
			$ vf-admin address remove 20
	`),
	Args: cobra.ExactArgs(1),
	RunE: remove.CmdRunE,
}

func init() {
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
