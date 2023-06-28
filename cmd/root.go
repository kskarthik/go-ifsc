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

// embed the IFSC.csv file into the binary
//
//go:embed IFSC.csv
var IFSCCodes string

// the parsed csv
var CsvSlice [][]string

// the column names of the csv
var Fields = []string{}

// ifsc codes as a map
var IFSCMap = make(map[string][]string)

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
	// create a map having ifsc code as key and of value slice
	for _, v := range CsvSlice[1:] {
		IFSCMap[v[1]] = v
	}
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

// checks whether a given IFSC code is valid, retuns a slice
func CheckIfSC(code string) ([]string, error) {

	var e error = errors.New("Record not found")
	// open bleve index
	index, _ := bleve.Open(IndexDir)
	// define a new query
	query := bleve.NewMatchQuery(strings.TrimSpace(code))
	searchRequest := bleve.NewSearchRequest(query)
	// enable all fields of the resulting document
	searchRequest.Fields = []string{"*"}
	result, _ := index.Search(searchRequest)
	// handle the case of no matching
	if result.Hits.Len() == 0 {
		return []string{}, e
	}
	//TODO: convert the result []interface{} to []string
	for _, val := range result.Hits[0].Fields {
		r := fmt.Sprintf("%s", val)
		print(r)
	}
	return []string{code}, e
}

// search the csv records which include the given search term
func SearchIFSC(searchTerm string) ([][]string, error) {
	c := make(chan [][]string)
	// trim the white spaces of the searchTerm if any
	keyWord := strings.TrimSpace(searchTerm)
	// check if search term is a valid ifsc code
	v, err := CheckIfSC(keyWord)
	if err == nil {
		return [][]string{v}, nil
	}
	// else create go routines to concurrenly search for
	// given input in different ranges of CsvSlice
	go searchSlice(keyWord, CsvSlice[1:50000], c)
	go searchSlice(keyWord, CsvSlice[50000:100000], c)
	go searchSlice(keyWord, CsvSlice[100000:], c)
	// assign goroutine results to three variables
	r1, r2, r3 := <-c, <-c, <-c
	// this var is used as return value
	var result [][]string
	// loop over all goroutine results and append to final result slice
	for _, s := range r1 {
		r := append(result, s)
		result = r
	}
	for _, s := range r2 {
		r := append(result, s)
		result = r
	}
	for _, s := range r3 {
		r := append(result, s)
		result = r
	}
	return result, nil
}

// this function is used as a goroutine
func searchSlice(keyWord string, slice [][]string, c chan [][]string) {
	searchResults := [][]string{}
	// loop over the csv fields
	for _, record := range slice {
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
				searchResults = append(searchResults, record)
				break
			}
		}
	}
	c <- searchResults
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
