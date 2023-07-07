/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: AGPLv3 (https://www.gnu.org/licenses/agpl-3.0.txt)

This file contains global vars and utility functions
*/
package cmd

import (
	"fmt"
	"strconv"
)

// this var stores the location of the bleve's index directory
var IndexDir string

var AppVersion string = "0.2.0"

// the columns of the csv
var Fields = [16]string{"BANK", "IFSC", "BRANCH", "CENTRE", "DISTRICT", "STATE", "ADDRESS", "CONTACT", "IMPS", "RTGS", "CITY", "ISO3166", "NEFT", "MICR", "UPI", "SWIFT"}

// debug or release mode for the rest api server
var ServerMode string

// default search result max limit
const DefaultSearchLimit int = 100

// default text matching pattern
const DefaultMatch = "fuzzy"

const SearchHelp string = `Text matching type:
	all - Matches docs containing all search termsany
	any - Matches docs containing any one of the search terms
	fuzzy - Matches docs containing any or similar search terms
	regex - Advanced query syntax. Refer https://blevesearch.com/docs/Query-String-Query/`

type SearchParams struct {
	// search terms
	terms []string
	// text matching pattern
	match string
	// result count limit
	limit int
}

// the body json result
type Body struct {
	BANK     string
	IFSC     string
	BRANCH   string
	CENTRE   string
	DISTRICT string
	STATE    string
	ADDRESS  string
	CONTACT  string
	IMPS     bool
	RTGS     bool
	CITY     string
	ISO3166  string
	NEFT     bool
	MICR     string
	UPI      bool
	SWIFT    *string
}

// convert a string to boolean
func ToBool(s string) bool {
	r, _ := strconv.ParseBool(s)
	return r
}

// converts the []interface{} to []string
func ConvertToSlice(fields map[string]interface{}) []string {

	var result []string

	for _, val := range fields {
		for _, v := range val.([]any) {
			result = append(result, v.(string))
		}
	}
	return result
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
		fmt.Printf("%8s\t%s\n", Fields[i], value)
	}
}