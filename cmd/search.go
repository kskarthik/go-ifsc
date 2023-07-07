/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var params SearchParams

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
	searchCmd.Flags().StringVarP(&params.match, "match", "m", DefaultMatch, SearchHelp)
	// search count limit
	searchCmd.Flags().IntVarP(&params.limit, "limit", "l", DefaultSearchLimit, "Limit the number of search results")
}
