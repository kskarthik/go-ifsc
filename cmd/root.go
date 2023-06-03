/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ifsc",
	Short: "Search & Validate IFSC Codes",
	Long:  `This utility helps to search, validate IFSC codes of Indian banks`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// embed the IFSC.csv file into the binary
//
//go:embed IFSC.csv
var IFSCCodes string

// the parsed csv
var CsvSlice [][]string

// the column names of the csv
var Fields = []string{}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ifsc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// read the csv
	r := csv.NewReader(strings.NewReader(IFSCCodes))
	// reads the string as a slice
	slice, readErr := r.ReadAll()
	if readErr != nil {
		return
	}
	// assign the csv slice & fields to respective global variables
	CsvSlice = slice
	Fields = CsvSlice[0]
}

/*
checks whether a given IFSC code is valid, retuns a slice

TODO: optimize the speed of validation, currenly using the linear approach
*/
func CheckIfSC(code string) ([]string, error) {
	// custom error
	var e error = errors.New("Record not found")
	// trim the white spaces for param
	c := strings.TrimSpace(code)
	// loop over the csv fields
	for _, record := range CsvSlice {
		// if code matches the record, return the result
		if c == record[1] {
			return record, nil
		}
	}
	return []string{code}, e
}

/*Print a search result to stdout*/
func PrintResult(record []string) {

	for i := range record {
		var value string = record[i]
		if record[i] == "true" {
			value = "yes"
		}
		if record[i] == "false" {
			value = "no"
		}
		if record[i] == "" {
			value = "N/A"
		}
		fmt.Println(Fields[i], ":", value)
	}
}

/*
search the csv records which include the given search term

TODO: optimize the search speed. Currenly using the linear search
Also, improve the handling of search params, The current accepts the
search param via cli argument & we the term has to be wrapped in quotes
for more than one word. eg "main road"
*/
func SearchIFSC(searchTerm string) ([][]string, error) {
	searchResults := [][]string{}
	// trim the white spaces of the searchTerm if any
	keyWord := strings.TrimSpace(searchTerm)
	// loop over the csv fields

	for _, record := range CsvSlice {
		// loop over all fields of a record
		for i := range record {
			// if exact word match is found
			if keyWord == record[i] {
				searchResults = append(searchResults, record)
				break
			}
			// if the search term matches any of the fields of the record
			if strings.Contains(strings.ToLower(record[i]), strings.ToLower(keyWord)) {
				// if found, append the record to the searchResults slice
				s := append(searchResults, record)
				searchResults = s
				break
			}
		}
	}
	return searchResults, nil
}
