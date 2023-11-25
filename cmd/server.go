/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: AGPLv3 (https://www.gnu.org/licenses/agpl-3.0.txt)
*/
package cmd

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
	"strings"
)

// cli falgs
var hostPort string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Launch the REST API server",
	Long: `Launch the REST API server at port 9000. Can be customized with flags

Endpoints:
	/ - validate a IFSC code. Returns an json object on success, else throws 404 error
	url params
	==========
	ifsc: valid ifsc code (string)

	/search - Search for banks / ifsc codes. returns an array of objects
	query params
	============
	q: search terms (string)
	match: can be of type all, any, fuzzy, adv (string)
	limit: limit of search results (int)`,

	Run: func(cmd *cobra.Command, args []string) {

		IndexDirExists()
		// server mode
		gin.SetMode(ServerMode)
		startServer()
	},
}

// convert & return the IFSC slice as struct
func ifscStruct(r []string) Body {

	var b Body

	b.BANK = r[0]
	b.IFSC = r[1]
	b.BRANCH = r[2]
	b.CENTRE = r[3]
	b.DISTRICT = r[4]
	b.STATE = r[5]
	b.ADDRESS = r[6]
	b.CONTACT = r[7]
	b.IMPS = ToBool(r[8])
	b.RTGS = ToBool(r[9])
	b.CITY = r[10]
	b.ISO3166 = r[11]
	b.NEFT = ToBool(r[12])
	b.MICR = r[13]
	b.UPI = ToBool(r[14])
	// return null of swift is empty
	if r[15] == "" {
		b.SWIFT = nil
	} else {
		b.SWIFT = &r[15]
	}
	return b
}

// start the REST api server, handle the config & incoming requests
func startServer() {
	router := gin.Default()

	// use defaults of cors package
	// https://github.com/gin-contrib/cors#using-defaultconfig-as-start-point
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	hello(router)
	fields(router)
	checkAPI(router)
	searchAPI(router)

	fmt.Printf("Starting server on http://0.0.0.0:%s\nPress Ctrl+C to stop\n", hostPort)

	// start the server
	router.Run(":" + hostPort)
}

// returns bank fields
func fields(router *gin.Engine) {

	router.GET("/fields", func(c *gin.Context) {
		c.JSON(http.StatusOK, Fields)
	})
	return
}

// returns app version
func hello(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"version": AppVersion})
	})
	return
}

// validate IFSC code, Throw 404 if invalid
func checkAPI(router *gin.Engine) {

	router.GET("/:ifsc", func(c *gin.Context) {

		ifscCode := c.Param("ifsc")

		if len(ifscCode) != 11 {
			c.Status(http.StatusNotFound)
			return
		} else {
			res, e := CheckIFSC(ifscCode)
			if e != nil {
				c.Status(http.StatusNotFound)
			} else {
				c.JSON(http.StatusOK, ifscStruct(res))
			}
		}
	})
}

// search for banks
func searchAPI(router *gin.Engine) {

	router.GET("/search", func(c *gin.Context) {
		// parse the query params
		params := SearchParams{}
		params.match = c.Query("match")

		// process the limit
		if c.Query("limit") == "" {
			params.limit = DefaultSearchLimit
		} else {
			r, _ := strconv.Atoi(c.Query("limit"))
			params.limit = r
		}
		// process the match
		if params.match == "" {
			params.match = DefaultMatch
		}
		// handle advanced search query
		if params.match != "adv" {
			params.terms = strings.Split(c.Query("q"), " ")
		} else {
			params.terms = []string{c.Query("q")}
		}

		var statusCode int

		response := []Body{}
		// perform the search
		res, e := params.SearchBanks()
		if e != nil {
			statusCode = http.StatusBadRequest
			return
		} else {
			// loop over the elements of the slice
			for i := range res {
				r := append(response, ifscStruct(res[i]))
				response = r
				statusCode = http.StatusOK
			}
		}
		c.JSON(statusCode, response)
	})
	return
}

func init() {
	rootCmd.AddCommand(serverCmd)
	// parse the cli flags
	serverCmd.Flags().StringVarP(&hostPort, "port", "p", "9000", "server port")
	serverCmd.Flags().StringVarP(&ServerMode, "mode", "m", "release", "Toggle between debug & release mode for server")
}
