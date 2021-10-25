package utils

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(writer io.Writer, headers []string, data [][]string) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader(headers)
	tl := len(headers)
	for _, v := range data {
		dl := len(v)
		if dl > tl {
			dl = tl
		}
		table.Append(v[:dl])
	}

	table.Render()
}
