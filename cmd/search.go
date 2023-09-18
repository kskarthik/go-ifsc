/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var params SearchParams
var output string
var columns []string

var columnHelp string = fmt.Sprintf("Customize the table columns. Available Columns: %s", strings.Join(Fields, ", "))

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for banks / IFSC codes",
	Long: `Search for banks / IFSC codes

	The search term can be anything: bank address, name, phone number, city etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide search term(s)")
			os.Exit(1)
		}
		params.terms = args
		// print the search results if there are any, to stdout
		searchResults, e := SearchIFSC(params)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		if len(searchResults) > 0 {
			columnIndexes := columnMap(columns)
			PrintTable(searchResults, columnIndexes)
		}
	},
}

// parse column names from cli and
// map them to corresponding index
func columnMap(v []string) (i []int) {

	// range over user defined columns
	for _, userColumn := range v {
		// range over all columns & check if user defined column name exists
		// if found, append the index & name to result
		for index, columnName := range Fields {
			if userColumn == columnName {
				i = append(i, index)
			}
		}
	}
	return
}

// print the result as cli table
func PrintTable(v [][]string, col []int) {

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	// config of columns
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     "ADDRESS",
			WidthMax: 25,
		},
		{
			Name:     "BANK",
			WidthMax: 18,
		},
		{
			Name:     "CITY",
			WidthMax: 15,
		},

		{
			Name:     "BRANCH",
			WidthMax: 15,
		},
	})
	// set table headers
	var tableRows table.Row
	// loop over all user provided colums- and set the table column names
	for _, val := range columns {
		tableRows = append(tableRows, val)
	}

	t.AppendHeader(tableRows)

	// for each search result, filter out the user selected fields
	for _, csv := range v {
		// create a temporary row
		var tmpRow table.Row
		// loop over all user defined table column indexes and
		// append the result to a temporary row
		for _, i := range col {
			tmpRow = append(tmpRow, csv[i])
		}
		t.AppendRow(tmpRow)
	}

	// render the output format
	switch output {
	case "html":
		t.RenderHTML()
	case "csv":
		t.RenderCSV()
	case "md":
		t.RenderMarkdown()
	default:
		// ascii
		t.Render()
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
	// initialize flags
	// matching pattern
	searchCmd.Flags().StringVarP(&params.match, "match", "m", DefaultMatch, SearchHelp)
	// search count limit
	searchCmd.Flags().IntVarP(&params.limit, "limit", "l", DefaultSearchLimit, "Limit the number of search results")
	// table output
	searchCmd.Flags().StringVarP(&output, "output", "o", "ascii", "specify output format, Available formats: html, md, csv")
	// customize table columns
	searchCmd.Flags().StringSliceVarP(&columns, "columns", "c", DefaultColumns, columnHelp)
}
