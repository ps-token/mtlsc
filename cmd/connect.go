package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"

	"github.com/spf13/cobra"
)

var (
	count       int
	concurrency int
	cert        string
	key         string
	server      []string
	wg          sync.WaitGroup
	c           chan struct{}
	successes   int
)

func certificatePool(server []string) *x509.CertPool {
	caCertPool := x509.NewCertPool()
	for _, v := range server {
		serverCert, err := ioutil.ReadFile(v)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool.AppendCertsFromPEM(serverCert)
	}
	return caCertPool
}

func clientCertificate(cert, key string) tls.Certificate {
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}
	return certificate
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a host and test mTLS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if concurrency == 0 || count == 0 {
			fmt.Println("[ERR] Cannot have 0 goroutines or 0 requests.")
			os.Exit(1)
		}

		uri := args[0]

		fmt.Printf("Testing mTLS to %s\n\n", uri)

		// Read client certificate and key
		certificate := clientCertificate(cert, key)
		fmt.Println("[OK] Read client certificate and key")

		pool := certificatePool(server)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: pool,
					// Certificates: []tls.Certificate{certificate},
					GetClientCertificate: func(_ *tls.CertificateRequestInfo) (*tls.Certificate, error) {
						return &certificate, nil
					},
					Renegotiation:      tls.RenegotiateOnceAsClient,
					InsecureSkipVerify: insecure,
				},
			},
		}
		fmt.Println("[OK] Initialized HTTP Client")

		fmt.Printf("\n[OK] Initiating test [%d connections | %d tests per connection]\n\n", concurrency, count)

		wg.Add(concurrency)
		c = make(chan struct{}, concurrency)
		for i := 0; i < concurrency; i++ {
			if debug {
				fmt.Printf("[DEBUG] Concurrency loop %d\n", i)
			}
			go func() {
				for range c {
					r, err := client.Get(uri)
					if err != nil {
						fmt.Printf("[ERR] Response => [HTTP %d %s]\n", r.StatusCode, http.StatusText(r.StatusCode))
						return
					}
					// force response body closure to ensure session reuse.
					io.Copy(ioutil.Discard, r.Body)
					r.Body.Close()
					fmt.Printf("[OK] Response => [HTTP %d %s]\n", r.StatusCode, http.StatusText(r.StatusCode))
					successes++
				}
				defer wg.Done()
			}()
		}

		requests := concurrency * count
		for i := 0; i < requests; i++ {
			if debug {
				fmt.Printf("[DEBUG] Request loop %d\n", i)
			}
			c <- struct{}{}
		}
		close(c)
		wg.Wait()

		fmt.Printf("\n[OK] %d/%d tests passed\n", successes, concurrency*count)
		fmt.Println("[OK] Test complete")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().IntVarP(&count, "count", "", 5, "The amount of connections to execute during the test")
	connectCmd.Flags().IntVarP(&concurrency, "concurrency", "", runtime.NumCPU()/2, "The number of concurrent connections to use for the test")
	connectCmd.Flags().StringSliceVarP(&server, "server", "", []string{}, "Path(s) to server side bundle or CA cert, intermediataries and leaf certificates (required)")
	connectCmd.Flags().StringVarP(&cert, "cert", "", "", "Path to your client-side certificate (required)")
	connectCmd.Flags().StringVarP(&key, "key", "", "", "Path to your client-side key (required)")

	// Mark flags as required
	connectCmd.MarkFlagRequired("server")
	connectCmd.MarkFlagRequired("cert")
	connectCmd.MarkFlagRequired("key")
}
