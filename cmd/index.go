/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strconv"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index the IFSC data locally",
	Long: `Index the IFSC data locally

	Downloads the latest IFSC csv dump & store in the system cache.
	This command should be invoked in two scenarios:

		1. During first setup
		2. When there is an upstream release of new IFSC dataset

	If XDG_CACHE_HOME env is set, it will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		getIFSCRelease()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}

// utility function
func httpGet(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		b, _ := io.ReadAll(response.Body)
		if err != nil {
			os.Exit(1)
		}
		return b
	}
	return []byte{}
}

// get the latest ifsc release info
func getIFSCRelease() {

	fmt.Println("Looking up latest IFSC release from https://github.com/razorpay/ifsc/releases")

	b := httpGet("https://api.github.com/repos/razorpay/ifsc/releases?per_page=1")

	body := []map[string]any{}

	err := json.Unmarshal(b, &body)
	if err != nil {
		os.Exit(1)
	}
	v := body[0]["name"]
	fmt.Println("Downloading IFSC.csv", v)
	url := fmt.Sprintf("https://github.com/razorpay/ifsc/releases/download/%s/IFSC.csv", v)
	res := httpGet(url)
	indexCSV(res)

}

// create an new bleve index using the csv values
func indexCSV(v []byte) {
	// convert byte to io.Reader which the csv reader requires
	r := bytes.NewReader(v)
	// invoke the csv reader
	s := csv.NewReader(r)
	// read csv
	csvSlice, _ := s.ReadAll()
	// delete the index dir, if already exists
	os.RemoveAll(IndexDir)
	// create a new index in user's cache directory
	index, newIndexErr := bleve.New(IndexDir, bleve.NewIndexMapping())
	if newIndexErr != nil {
		fmt.Println(newIndexErr)
		os.Exit(1)
	}
	fmt.Printf("Indexing the data in '%s' This might take a few minutes\n", IndexDir)

	// invoke bleve's batch indexing https://github.com/blevesearch/bleve/discussions/1834#discussioncomment-6280490
	batch := index.NewBatch()

	for i := range csvSlice[1:] {
		batch.Index(strconv.Itoa(i), csvSlice[i])
	}
	// append the created batch to index
	indexingErr := index.Batch(batch)

	if indexingErr != nil {
		fmt.Println(indexingErr)
		os.Exit(1)
	}
	fmt.Println("Indexing Complete!")
}
