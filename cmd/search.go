/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Fuzzy search for banks / IFSC codes",
	Long: `Fuzzy search for banks / IFSC codes

	The search term can be anything, Like Bank address, name, phone number, location etc`,
	Run: func(cmd *cobra.Command, args []string) {
		var searchString string
		// convert all additional arguments to string
		for i := range args {
			searchString += " " + args[i]
		}
		// print the search results if there are any, to stdout
		searchResults, e := SearchIFSC(searchString)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
