/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"github.com/spf13/cobra"
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ifsc",
	Short: "Search & Validate IFSC Codes",
	Long: `This utility helps to search, validate IFSC codes of Indian banks`,
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ifsc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// embed the IFSC.csv file into the binary
//
//go:embed IFSC.csv
var IFSCCodes string

// IFSC fields
var Fields = [16]string{"BANK", "IFSC", "BRANCH", "CENTRE", "DISTRICT", "STATE", "ADDRESS", "CONTACT", "IMPS", "RTGS", "CITY", "ISO3166", "NEFT", "MICR", "UPI", "SWIFT"}

/* checks whether a given IFSC code is valid, retuns a slice
TODO:optimize the speed of validation, currenly using the linear approach
*/
func CheckIfSC(code []string) ([]string, error) {
	// read the csv
	r := csv.NewReader(strings.NewReader(IFSCCodes))
	var e error = errors.New("Record not found")
	// loop over the csv fields
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// if code matches the record, return the result
		if code[0] == record[1] {
			return record, nil
		}
	}
	return code, e
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
			value = "?"
		}
		fmt.Println(Fields[i], ":", value)
	}
}

/* search the csv records which include the given search term
TODO: optimize the search speed. Currenly using the linear search
Also, improve the handling of search params, The current accepts the 
search param via cli argument & we the term has to be wrapped in quotes
for more than one word. eg "main road"
*/
func SearchIFSC(searchTerm string) ([][]string, error) {
	// read the csv
	r := csv.NewReader(strings.NewReader(IFSCCodes))
	searchResults := [][]string{}
	// loop over the csv fields
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// loop over all fields of a record
		for i := range record {
			// convert the strings to lower case & compare
			// if the search term matches any of the fields of the record
			if strings.Contains(strings.ToLower(record[i]), strings.ToLower(searchTerm)) {
				// if found, append the record to the searchResults slice
				s := append(searchResults, record)
				searchResults = s
				break
			}
		}
	}
	return searchResults, nil
}
