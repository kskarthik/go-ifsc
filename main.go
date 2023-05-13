package main

import (
  "ifsc/cmd"
)

func main() {
  cmd.Execute()
}

/*
Author: <kskarthik@disroot.org>
License: GPLv3
*/
// package main
//
// import (
// 	_ "embed"
// 	"encoding/csv"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )
//
// // embed the IFSC.csv file into the binary
// //
// //go:embed IFSC.csv
// var ifscCodes string
//
// // IFSC fields
// var fields = [16]string{"BANK", "IFSC", "BRANCH", "CENTRE", "DISTRICT", "STATE", "ADDRESS", "CONTACT", "IMPS", "RTGS", "CITY", "ISO3166", "NEFT", "MICR", "UPI", "SWIFT"}
//
// var help = map[string]string{
// 	"arg": "No arguments specified",
// 	"all": "This utility shows the bank details of given IFSC code\n\n USAGE: ./ifsc [COMMAND] [INPUT]  \n\n COMMANDS: \n\tcheck - checks the given IFSC code & return the bank details if valid\n\tsearch - return results of banks based on keyword\n\tserve - starts the REST API server [TODO]"}
//
// func main() {
// 	// parse the cli arguments
// 	args := os.Args[1:]
// 	// if no argument is given handle the case
// 	if len(args) == 0 {
// 		fmt.Println(help["all"])
// 		return
// 	}
// 	if args[0] == "check" && len(args) == 2 {
// 		result, err := checkIfscCode(args[1])
// 		if err == nil {
// 			printResult(result)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	} else if args[0] == "search" {
// 		searchResults, e := searchIFSC(args[1])
// 		// print the search results if there are any, to stdout
// 		if e == nil && len(searchResults) > 0 {
// 			for i := 0; i < len(searchResults); i++ {
// 				printResult(searchResults[i])
// 				fmt.Println("----------------------")
// 			}
// 			// display the result count after the last result
// 			fmt.Println(len(searchResults), "results")
// 		}
//
// 	} else {
// 		fmt.Println(help["all"])
// 	}
// }
//
// /* checks whether a given IFSC code is valid, retuns a slice
// TODO:optimize the speed of validation, currenly using the linear approach
// */
// func checkIfscCode(code string) ([]string, error) {
// 	// read the csv
// 	r := csv.NewReader(strings.NewReader(ifscCodes))
// 	var e error = errors.New("Record not found")
// 	// loop over the csv fields
// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// if code matches the record, return the result
// 		if code == record[1] {
// 			return record, nil
// 		}
// 	}
// 	return []string{code}, e
// }
//
// /*Print a search result to stdout*/
// func printResult(record []string) {
//
// 	for i := range record {
// 		var value string = record[i]
// 		if record[i] == "true" {
// 			value = "yes"
// 		}
// 		if record[i] == "false" {
// 			value = "no"
// 		}
// 		if record[i] == "" {
// 			value = "?"
// 		}
// 		fmt.Println(fields[i], ":", value)
// 	}
// }
//
// /* search the csv records which include the given search term
// TODO: optimize the search speed. Currenly using the linear search
// Also, improve the handling of search params, The current accepts the 
// search param via cli argument & we the term has to be wrapped in quotes
// for more than one word. eg "main road"
// */
// func searchIFSC(searchTerm string) ([][]string, error) {
// 	// read the csv
// 	r := csv.NewReader(strings.NewReader(ifscCodes))
// 	searchResults := [][]string{}
// 	// loop over the csv fields
// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// loop over all fields of a record
// 		for i := range record {
// 			// convert the strings to lower case & compare
// 			// if the search term matches any of the fields of the record
// 			if strings.Contains(strings.ToLower(record[i]), strings.ToLower(searchTerm)) {
// 				// if found, append the record to the searchResults slice
// 				s := append(searchResults, record)
// 				searchResults = s
// 			}
// 		}
// 	}
// 	return searchResults, nil
// }
