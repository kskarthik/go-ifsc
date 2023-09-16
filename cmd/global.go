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

var AppVersion string = "0.3.1"

// the columns of the csv
var Fields = []string{"BANK", "IFSC", "BRANCH", "CENTRE", "DISTRICT", "STATE", "ADDRESS", "CONTACT", "IMPS", "RTGS", "CITY", "ISO3166", "NEFT", "MICR", "UPI", "SWIFT"}

// default columns appearing during CLI search
var DefaultColumns []string = []string{"BANK", "BRANCH", "ADDRESS", "IFSC"}

// debug or release mode for the rest api server
var ServerMode string

// default search result max limit
const DefaultSearchLimit int = 20

// default text matching pattern
const DefaultMatch = "any"

const SearchHelp string = `Text matching options:
	all - Matches docs containing all search terms
	any - Matches docs containing any one of the search terms
	fuzzy - Matches docs containing any one or similar search terms
	adv - Advanced query syntax. Refer https://blevesearch.com/docs/Query-String-Query/`

type SearchParams struct {
	// search terms provided by the user
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

		value := record[i]

		switch value {
		case "true":
			value = "yes"
		case "false":
			value = "no"
		case "":
			value = "N/A"
		}

		fmt.Printf("%8s\t%s\n", Fields[i], value)
	}
}
