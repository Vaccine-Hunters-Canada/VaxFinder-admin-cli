package locationutils

import (
	"fmt"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// ConvertExpandedLocationJSONToTableRow returns a list of strings for the table row of an expanded location JSON response
func ConvertExpandedLocationJSONToTableRow(json *api.LocationExpandedResponse) []string {
	data := []string{
		strconv.Itoa(json.Id),
		json.Name,
		strconv.FormatBool(json.Active != 0),
		utils.CoalesceString(json.Postcode),
		utils.CoalesceString(json.Phone),
		utils.CoalesceString(json.Notes),
		utils.CoalesceString(json.Url),
		utils.CoalesceString(json.Tags),
	}

	orgString := ""
	if json.Organization != nil {
		orgString = fmt.Sprintf("organization id: %d\n%s", json.Organization.Id, utils.CoalesceString(&json.Organization.ShortName))
	}
	data = append(data, orgString)

	addrString := ""
	if json.Address != nil {
		addrString = fmt.Sprintf("address id: %d\n%s\n%s\n%s\n%s",
			json.Address.Id, utils.CoalesceString(json.Address.Line1),
			utils.CoalesceString(json.Address.City),
			utils.CoalesceString(&json.Address.Province),
			utils.CoalesceString(&json.Address.Postcode),
		)
	}

	data = append(data, addrString)

	data = append(data, json.CreatedAt.String())

	return data
}

// ConvertLocationJSONToTableRow returns a list of strings for the table row of a location JSON response
func ConvertLocationJSONToTableRow(json *api.LocationResponse) []string {
	data := []string{
		strconv.Itoa(json.Id),
		json.Name,
		strconv.FormatBool(json.Active != 0),
		utils.CoalesceString(json.Postcode),
		utils.CoalesceString(json.Phone),
		utils.CoalesceString(json.Notes),
		utils.CoalesceString(json.Url),
		utils.CoalesceString(json.Tags),
	}

	orgString := ""
	if json.Organization != nil {
		orgString = fmt.Sprintf("org id: %d", *json.Organization)
	}
	data = append(data, orgString)

	addrString := ""

	if json.Address != nil {
		addrString = fmt.Sprintf("address id: %d", *json.Address)
	}

	data = append(data, addrString)

	data = append(data, json.CreatedAt.String())

	return data
}
