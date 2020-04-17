/*
Copyright © 2020 Token Inc <ops@token.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// connecturiCmd represents the connecturi command
var connecturiCmd = &cobra.Command{
	Use:   "connecturi",
	Short: "The URI to connect to",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connecturi called")
	},
}

func init() {
	connectCmd.AddCommand(connecturiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connecturiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connecturiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
