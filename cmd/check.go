/*
Copyright © 2023 Sai Karthik <kskarthik@disroot.org>
License: GPLv3
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whether a given IFSC code is valid",
	Long: `Check whether a given IFSC code is valid

	if the given IFSC code is valid, the details of the bank is returned
	Incorrect / invalid IFSC code will return error`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args[0]) != 11 {
			fmt.Println("Invalid IFSC Code")
			os.Exit(1)
		}
		r, err := CheckIfSC(args[0])
		if err == nil {
			PrintResult(r)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
