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

// ifsc codes as a map
var IFSCMap = make(map[string][]string)

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
	// create a map having ifsc code as key and of value slice
	for _, v := range CsvSlice[1:] {
		IFSCMap[v[1]] = v
	}
}

// checks whether a given IFSC code is valid, retuns a slice
func CheckIfSC(code string) ([]string, error) {
	// custom error
	var e error = errors.New("Record not found")
	// trim the white spaces for param
	c := strings.TrimSpace(code)
	// if the key exists in the map, return it's value
	// else throw err
	val, exists := IFSCMap[c]
	if exists {
		return val, nil
	}
	return []string{code}, e
}

/* func CheckIfSC(code string) ([]string, error) {
	// custom error
	var e error = errors.New("Record not found")
	// trim the white spaces for param
	ifscCode := strings.TrimSpace(code)
	// create a channel
	c := make(chan []string)
	// create go routines to concurrenly search for
	// given input in different ranges of CsvSlice
	go checkSlice(ifscCode, CsvSlice[1:50000], c)
	go checkSlice(ifscCode, CsvSlice[50000:100000], c)
	go checkSlice(ifscCode, CsvSlice[100000:], c)
	// assign goroutine results to three variables
	r1, r2, r3 := <-c, <-c, <-c
	// check each result & return the one which has the value
	if len(r1) != 0 {
		return r1, nil
	}

	if len(r2) != 0 {
		return r2, nil
	}

	if len(r3) != 0 {
		return r3, nil
	}
	return []string{code}, e
} */

// loop over the csv fields & return the matching result to channel
/* func checkSlice(input string, slice [][]string, c chan []string) {
	var result []string
	for _, record := range slice {
		// if code matches the record, return the result
		if input == record[1] {
			result = record
		}
	}
	c <- result
} */

// search the csv records which include the given search term
func SearchIFSC(searchTerm string) ([][]string, error) {
	c := make(chan [][]string)
	// trim the white spaces of the searchTerm if any
	keyWord := strings.TrimSpace(searchTerm)
	// if the search term is a vaild ifsc code, return it's data
	v, err := CheckIfSC(keyWord)
	if err != nil {
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
