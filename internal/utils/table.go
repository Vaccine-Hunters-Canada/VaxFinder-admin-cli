package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

// RenderDefaultTable renders the default ASCII table on stdout
func RenderDefaultTable(colNames []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(colNames)
	table.SetBorder(false)

	for _, v := range data {
		table.Append(v)
	}

	fmt.Print("\n")
	table.Render()
}
