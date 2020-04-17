/*
Copyright © 2020 Token Inc <ops@token.io>

*/
package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	server []string
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a host and test mTLS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		uri := args[0]
		count, _ := cmd.Flags().GetInt("count")
		fmt.Printf("Connecting to %s [%d connections]\n", uri, count)
		fmt.Println()

		// Read client certificate and key
		certificate, _ := cmd.Flags().GetString("cert")
		key, _ := cmd.Flags().GetString("key")
		cert, err := tls.LoadX509KeyPair(certificate, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[OK] Read client certificate and key")

		// Read server certs and add to a cert pool
		fmt.Println("[OK] Created Certificate Pool")
		caCertPool := x509.NewCertPool()
		for _, v := range server {
			serverCert, err := ioutil.ReadFile(v)
			if err != nil {
				log.Fatal(err)
			}
			caCertPool.AppendCertsFromPEM(serverCert)
			fmt.Printf("[OK] Added %s to certificate pool\n", v)
		}

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      caCertPool,
					Certificates: []tls.Certificate{cert},
				},
			},
		}

		fmt.Println("[OK] Initialized HTTP Client")
		fmt.Println()

		for i := 0; i <= count; i++ {
			r, err := client.Get(uri)
			if err != nil {
				fmt.Printf("[ERR] Could not connect to %s [HTTP %d %s]\n", uri, r.StatusCode, http.StatusText(r.StatusCode))
				return
			}
			io.Copy(ioutil.Discard, r.Body)
			r.Body.Close()
			fmt.Printf("[OK] Connected to %s [HTTP %d %s]\n", uri, r.StatusCode, http.StatusText(r.StatusCode))
		}

		fmt.Println("\n[OK] Test complete")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().Int("count", 10, "The amount of connections to execute during the test")
	connectCmd.Flags().StringSliceVarP(&server, "server", "", []string{}, "Path(s) to server side bundle or CA cert, intermediataries and leaf certificates (required)")
	connectCmd.Flags().String("cert", "", "Path to your client-side certificate (required)")
	connectCmd.Flags().String("key", "", "Path to your client-side key (required)")

	// Mark flags as required
	connectCmd.MarkFlagRequired("server")
	connectCmd.MarkFlagRequired("cert")
	connectCmd.MarkFlagRequired("key")
}
