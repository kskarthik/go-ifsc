/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/blevesearch/bleve/v2"
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

// the column names of the csv
var Fields = [16]string{"BANK", "IFSC", "BRANCH", "CENTRE", "DISTRICT", "STATE", "ADDRESS", "CONTACT", "IMPS", "RTGS", "CITY", "ISO3166", "NEFT", "MICR", "UPI", "SWIFT"}

// this var stores the location of the bleve's index directory
var IndexDir string

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ifsc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setCacheDir()
}

// Get user's cache dir.
// Respect the XDG env, if set
func setCacheDir() {

	xdgCachePath := os.Getenv("XDG_CACHE_HOME")

	if xdgCachePath != "" {
		IndexDir = xdgCachePath + "/ifsc"
		// fallback to default cache path
	} else {
		usrCacheDir, err := os.UserCacheDir()
		if err != nil {
			fmt.Println("Unable to locate cache directory")
			os.Exit(1)
		}
		IndexDir = usrCacheDir + "/ifsc"
	}
}

// convert the result []interface{} to []string
func convertToSlice(fields map[string]interface{}) []string {

	var result []string

	for _, val := range fields {
		for _, v := range val.([]any) {
			result = append(result, v.(string))
		}
	}
	return result
}

// checks whether a given IFSC code is valid, retuns a slice
func CheckIfSC(code string) ([]string, error) {

	var e error = errors.New("Record not found")
	// open bleve index
	index, _ := bleve.Open(IndexDir)
	defer index.Close()
	// define a new query
	query := bleve.NewMatchQuery(strings.TrimSpace(code))
	searchRequest := bleve.NewSearchRequest(query)
	// enable all fields of the resulting document
	searchRequest.Fields = []string{"*"}
	result, _ := index.Search(searchRequest)
	// handle the case of no matching
	if result.Hits.Len() == 0 {
		return []string{}, e
	} else {
		return convertToSlice(result.Hits[0].Fields), nil
	}
}

// search the csv records which include the given search term
func SearchIFSC(searchTerm string) ([][]string, error) {
	// open bleve index
	index, _ := bleve.Open(IndexDir)
	defer index.Close()
	// define a new query
	query := bleve.NewMatchQuery(strings.TrimSpace(searchTerm))
	searchRequest := bleve.NewSearchRequest(query)
	// enable all fields of the resulting document
	searchRequest.Fields = []string{"*"}
	result, _ := index.Search(searchRequest)
	// handle the case of no matching
	var finalResult [][]string
	// append the results to finalResult slice
	if result.Hits.Len() > 0 {
		for i := range result.Hits {
			finalResult = append(finalResult, convertToSlice(result.Hits[i].Fields))
		}
	}
	return finalResult, nil
}

// format the provided arg and print to stdout
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
