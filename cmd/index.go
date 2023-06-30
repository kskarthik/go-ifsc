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
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index IFSC data",
	Long: `Indexes the IFSC data locally

	Download the latest IFSC csv & store in cache.
	This command should be invoked before launching the ifsc command for the first time
	or to update to latest IFSC dataset with upstream releases
	`,
	Run: func(cmd *cobra.Command, args []string) {
		getIFSCRelease()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

	fmt.Println("Checking for the latest IFSC release from https://github.com/razorpay/ifsc/releases")

	b := httpGet("https://api.github.com/repos/razorpay/ifsc/releases?per_page=1")

	body := []map[string]any{}

	err := json.Unmarshal(b, &body)
	if err != nil {
		os.Exit(1)
	}
	v := body[0]["name"]
	fmt.Println("Downloading the IFSC.csv", v)
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
	fmt.Printf("Indexing the csv data in %s This may take a few minutes\n", IndexDir)
	// invoke batch indexing https://github.com/blevesearch/bleve/discussions/1834#discussioncomment-6280490
	batch := index.NewBatch()
	for i := range csvSlice {
		// do not index columns
		if i != 0 {
			batch.Index(csvSlice[i][1], csvSlice[i])
		}
	}
	// append the created batch to index
	indexingErr := index.Batch(batch)
	if indexingErr != nil {
		fmt.Println(indexingErr)
		os.Exit(1)
	}
	fmt.Println("Indexing Complete!")
}
