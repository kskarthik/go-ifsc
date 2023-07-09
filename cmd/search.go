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
)

var params SearchParams
var output string

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
			PrintTable(searchResults)
		}
	},
}

// print the result as cli table
func PrintTable(v [][]string) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"IFSC", "BANK", "CITY", "STATE", "ADDRESS"})
	// show result count in footer
	if len(v) > 3 {
		t.AppendFooter(table.Row{" ", " ", "", "", fmt.Sprintf("total: %v", len(v))})
	}
	for _, csv := range v {
		row := table.Row{csv[1], csv[0], csv[10], csv[5], csv[6]}
		t.AppendRow(row)
	}
	// config of columns
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     "ADDRESS",
			WidthMax: 25,
		},
		{
			Name:     "BANK",
			WidthMax: 15,
		},
		{
			Name:     "CITY",
			WidthMax: 15,
		},
	})
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
	// flags
	searchCmd.Flags().StringVarP(&params.match, "match", "m", DefaultMatch, SearchHelp)
	// search count limit
	searchCmd.Flags().IntVarP(&params.limit, "limit", "l", DefaultSearchLimit, "Limit the number of search results")
	searchCmd.Flags().StringVarP(&output, "output", "o", "ascii", "specify output format, Available formats: html, md, csv")
}
