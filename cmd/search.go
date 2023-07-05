/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var match string
var limit int

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
		var searchArgs SearchParams
		searchArgs.match = match
		searchArgs.limit = limit
		searchArgs.terms = args
		// print the search results if there are any, to stdout
		searchResults, e := SearchIFSC(searchArgs)
		if e == nil && len(searchResults) > 0 {
			for i := range searchResults {
				PrintResult(searchResults[i])
				fmt.Println("----------------------")
			}
			// display the result count after the last result
			fmt.Println(len(searchResults), "results")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	// flags
	searchCmd.Flags().StringVarP(&match, "match", "m", DefaultMatch, "Set text matching type:\n\tall - Matches docs containing all search terms\n\tany - Matches docs containing any one of the search terms\n\tfuzzy - Matches docs containing similar search terms")
	// search count limit
	searchCmd.Flags().IntVarP(&limit, "limit", "l", DefaultSearchLimit, "Limit the number of search results")
}
