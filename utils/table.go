package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func ReadOnTable(columnNames []string, resultList [][]interface{}) {

	// table writer
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)

	rowHeader := table.Row{}
	for _, col := range columnNames {
		rowHeader = append(rowHeader, col)
	}
	//AppendHeader from columnNames
	tw.AppendHeader(rowHeader)
	for _, row := range resultList {
		tw.AppendRow(row)
	}
	tw.Render()

}
