/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ifsc",
	Short:   "Search & Validate IFSC Codes",
	Long:    `This utility helps to search, validate IFSC codes of Indian banks`,
	Version: AppVersion,
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
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	locateIndexDir()
}

// locate user's cache dir.
// Respect the XDG env, if set
func locateIndexDir() {

	dirName := "/ifsc"

	xdgCachePath := os.Getenv("XDG_CACHE_HOME")

	if xdgCachePath != "" {
		IndexDir = xdgCachePath + dirName
		// fallback to default cache path
	} else {
		usrCacheDir, err := os.UserCacheDir()
		if err != nil {
			fmt.Println("Unable to locate cache directory")
			os.Exit(1)
		}
		IndexDir = usrCacheDir + dirName
	}
}

// checks whether a given IFSC code is valid, retuns a slice
func CheckIFSC(code string) ([]string, error) {
	// open bleve index
	index, err := bleve.Open(IndexDir)
	if err != nil {
		fmt.Printf("Index does not exist! Create one first\n\n")
		rootCmd.Help()
		os.Exit(1)
	}
	defer index.Close()

	var e error = errors.New("Record not found")
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
		return ConvertToSlice(result.Hits[0].Fields), nil
	}
}

// prepares the search query based on the params
func SearchIFSC(q SearchParams) ([][]string, error) {

	bq := []query.Query{}
	// convert the query terms based on the matching option to bleve query type
	for _, term := range q.terms {
		if q.match == "fuzzy" {
			bq = append(bq, bleve.NewFuzzyQuery(term))
		} else {
			bq = append(bq, bleve.NewMatchQuery(term))
		}

	}

	var result [][]string

	switch q.match {
	// https://blevesearch.com/docs/Query-String-Query/
	case "adv":
		r, e := search(bleve.NewQueryStringQuery(q.terms[0]), q)
		if e != nil {
			return result, e
		}
		result = r
	case "all":
		r, e := search(query.NewConjunctionQuery(bq), q)
		if e != nil {
			return result, e
		}
		result = r
	default:
		r, e := search(query.NewDisjunctionQuery(bq), q)
		if e != nil {
			return result, e
		}
		result = r
	}
	return result, nil
}

// performs the actual search over the docs
func search(q query.Query, p SearchParams) ([][]string, error) {
	// open bleve index
	index, err := bleve.Open(IndexDir)
	if err != nil {
		fmt.Printf("Index does not exist! Create one first\n\n")
		rootCmd.Help()
		os.Exit(1)
	}
	defer index.Close()

	searchRequest := bleve.NewSearchRequest(q)
	// enable all fields of the resulting document
	searchRequest.Fields = []string{"*"}
	// max count of search results
	searchRequest.Size = p.limit
	// assign the search results
	result, searchErr := index.Search(searchRequest)

	// contains the results slice
	var finalResult [][]string

	if searchErr != nil {
		return finalResult, searchErr
	}
	// append the results to finalResult slice
	if result.Hits.Len() > 0 {
		for i := range result.Hits {
			finalResult = append(finalResult, ConvertToSlice(result.Hits[i].Fields))
		}
	}
	return finalResult, nil
}
