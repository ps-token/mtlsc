package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	debug    bool
	insecure bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mtlsc",
	Short: "A tool for testing mTLS",
	Long: `Test mTLS against a server by passing the tool your client certificate
and key, as well as a CA cert and leaf cert for the server. The CA cert can take
the form of a chain of intermediate certs that lead to the root CA cert.

From there, the tool will establish a connection, and execute a number of connections
against the given URI.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Run in debug mode")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "k", false, "Bypass certificate verification")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
