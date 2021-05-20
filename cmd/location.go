package cmd

import (
	"vf-admin/internal/cmdrun/location/add"
	"vf-admin/internal/cmdrun/location/get"
	"vf-admin/internal/cmdrun/location/list"
	"vf-admin/internal/cmdrun/location/remove"
	"vf-admin/internal/cmdrun/location/update"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// Command: `vf-admin location`
var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Manage locations",
}

// Command: `vf-admin location get <id>`
var locationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a location with a specified id",
	Example: heredoc.Doc(`
			# Get the location with id 1.
			$ vf-admin location get 1
	`),
	Args: cobra.ExactArgs(1),
	RunE: get.CmdRunE,
}

// Command: `vf-admin location list`
var locationListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of locations",
	Args:  cobra.ExactArgs(0),
	Run:   list.CmdRun,
}

// Command: `vf-admin location add <id>`
var locationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new location",
	Example: heredoc.Doc(`
			# Add a new active location with name "Guelph Hospital", postal code "N1E 4J4", website URL "http://www.gghorg.ca/", phone "(519) 822-5350", notes "Please call ahead to make an appointment.", tags "Guelph, Appointment", organization id 23 and address id 352.
			$ vf-admin location add --active 1 --name "Guelph Hospital" --postcode "N1E4J4" --url "http://www.gghorg.ca" --phone "(519) 822-5350" --notes "Please call ahead to make an appointment." --tags "Guelph, Appointment" --organization 23 --address 352
	`),
	Args: cobra.ExactArgs(0),
	RunE: add.CmdRunE,
}

// Command: `vf-admin location update <id>`
var locationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a location with a specified id",
	Example: heredoc.Doc(`
	# Add the location with id 15 to be active, have name "Guelph Hospital", postal code "N1E 4J4", website URL "http://www.gghorg.ca/", phone "(519) 822-5350", notes "Please call ahead to make an appointment.", tags "Guelph, Appointment", organization id 23 and address id 352.
	$ vf-admin location update 15 --active 1 --name "Guelph Hospital" --postcode "N1E4J4" --url "http://www.gghorg.ca" --phone "(519) 822-5350" --notes "Please call ahead to make an appointment." --tags "Guelph, Appointment" --organization 23 --address 352
	`),
	Args: cobra.ExactArgs(1),
	RunE: update.CmdRunE,
}

// Command: `vf-admin location remove <id>`
var locationRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a location with a specified id",
	Example: heredoc.Doc(`
			# Remove the location with id 8.
			$ vf-admin location remove 8
	`),
	Args: cobra.ExactArgs(1),
	RunE: remove.CmdRunE,
}

func init() {
	RootCmd.AddCommand(locationCmd)

	// Get command
	locationCmd.AddCommand(locationGetCmd)

	// List command
	locationCmd.AddCommand(locationListCmd)

	// Add command
	locationAddCmd.Flags().String("name", "", "name of location")
	_ = locationAddCmd.MarkFlagRequired("name")
	locationAddCmd.Flags().Int("active", 1, "is this location active? 1 or 0")
	_ = locationAddCmd.MarkFlagRequired("active")
	locationAddCmd.Flags().String("postcode", "", "postal code of location")
	locationAddCmd.Flags().String("phone", "", "phone number of location")
	locationAddCmd.Flags().String("notes", "", "notes about location")
	locationAddCmd.Flags().String("url", "", "website URL of location")
	locationAddCmd.Flags().String("tags", "", "search tags of location")
	locationAddCmd.Flags().Int("organization", 0, "id of organization running location")
	locationAddCmd.Flags().Int("address", 0, "id of address of location")

	locationCmd.AddCommand(locationAddCmd)

	// // Update command
	locationUpdateCmd.Flags().String("name", "", "name of location")
	_ = locationUpdateCmd.MarkFlagRequired("name")
	locationUpdateCmd.Flags().Int("active", 1, "is this location active? 1 or 0")
	_ = locationUpdateCmd.MarkFlagRequired("active")
	locationUpdateCmd.Flags().String("postcode", "", "postal code of location")
	locationUpdateCmd.Flags().String("phone", "", "phone number of location")
	locationUpdateCmd.Flags().String("notes", "", "notes about location")
	locationUpdateCmd.Flags().String("url", "", "website URL of location")
	locationUpdateCmd.Flags().String("tags", "", "search tags of location")
	locationUpdateCmd.Flags().Int("organization", 0, "id of organization running location")
	locationUpdateCmd.Flags().Int("address", 0, "id of address of location")

	locationCmd.AddCommand(locationUpdateCmd)

	// Remove command
	locationCmd.AddCommand(locationRemoveCmd)
}
