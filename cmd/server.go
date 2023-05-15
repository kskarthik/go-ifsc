/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
)

// cli command flags
var hostPort string
var Mode string

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
	// TODO: make this type string or nil
	SWIFT string
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Launch the REST API server",
	Long: `Launch the REST API server at port 9000. Can be customized with flags

Endpoints:
	/:ifsc - validate a IFSC code. Returns an json object on success, else throws 404 error
	/search/:query - Search for banks / ifsc codes`,
	Run: func(cmd *cobra.Command, args []string) {
		// server mode
		gin.SetMode(Mode)
		startServer()
	},
}

// convert a string to boolean
func toBool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}

// convert & return the IFSC slice as struct
func ifscStruct(r []string) Body {

	var body Body

	body.BANK = r[0]
	body.IFSC = r[1]
	body.BRANCH = r[2]
	body.CENTRE = r[3]
	body.DISTRICT = r[4]
	body.STATE = r[5]
	body.ADDRESS = r[6]
	body.CONTACT = r[7]
	body.IMPS = toBool(r[8])
	body.RTGS = toBool(r[9])
	body.CITY = r[10]
	body.ISO3166 = r[11]
	body.NEFT = toBool(r[12])
	body.MICR = r[13]
	body.UPI = toBool(r[14])
	body.SWIFT = r[15]

	return body
}

// start the REST api server & handle the config & incoming requests
func startServer() {
	router := gin.Default()
	// validate IFSC code
	router.GET("/:ifsc", func(c *gin.Context) {
		name := c.Param("ifsc")
		if len(name) != 11 {
			c.Status(http.StatusNotFound)
		}
		res, e := CheckIfSC(name)
		if e != nil {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, ifscStruct(res))
		}
	})
	// search for banks
	router.GET("/search/:query", func(c *gin.Context) {
		name := c.Param("query")
		res, e := SearchIFSC(name)
		if e != nil {
			c.Status(http.StatusNotFound)
		}
		array := []Body{}
		// loop over the elements of the slice
		for i := range res {
			r := append(array, ifscStruct(res[i]))
			array = r
		}
		c.JSON(http.StatusOK, array)
	})
	fmt.Printf("Starting server on http://localhost:%s\nPress Ctrl+C to stop", hostPort)
	// start the server
	router.Run(":" + hostPort)
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringVarP(&hostPort, "port", "p", "9000", "server port")
	serverCmd.Flags().StringVarP(&Mode, "mode", "m", "release", "Toggle between debug & release mode for server. Debug prints the logs")
}
