package validation

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/miekg/dns"
)

type ActiveUrlValidation struct {
	Value interface{}
	Field string
}

// Validate performs the active_url validation.
func (auv *ActiveUrlValidation) Validate() error {
	value := auv.Value.(string)

	// Parse the URL to validate the format
	u, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("failed! Invalid URL: %s", value)
	}

	// Check if the parsed URL has a valid host
	if u.Host == "" {
		return fmt.Errorf("failed! Invalid hostname: %s", value)
	}

	// Extract the hostname from the URL
	hostname := u.Hostname()
	if hostname == "" {
		return fmt.Errorf("failed! Invalid hostname: %s", value)
	}

	// Perform DNS resolution to get the A or AAAA record
	_, err = lookupAorAAAA(hostname)
	if err != nil {
		return fmt.Errorf("failed! DNS lookup failed for hostname: %s", hostname)
	}

	// Make an HTTP GET request to the URL in a separate goroutine
	respCh := make(chan *http.Response)
	errCh := make(chan error)

	go func() {
		resp, err := http.Get(u.String())
		if err != nil {
			errCh <- fmt.Errorf("failed! HTTP request failed for URL: %s", value)
			return
		}
		respCh <- resp
	}()

	// Set a timeout for the HTTP request
	timeout := time.Second * 5

	select {
	case resp := <-respCh:
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic("failed! Unable to close HTTP response body on ActiveUrlValidation")
			}
		}(resp.Body)

		// Check the response status code
		if resp.StatusCode < 200 || resp.StatusCode > 399 {
			return fmt.Errorf("failed! Invalid URL: %s", value)
		}

	case err := <-errCh:
		return err

	case <-time.After(timeout):
		return fmt.Errorf("failed! HTTP request timed out for URL: %s", value)
	}

	return nil
}

// lookupAorAAAA performs DNS resolution and returns the A or AAAA record for the given hostname.
func lookupAorAAAA(hostname string) ([]string, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(hostname), dns.TypeA)
	m.MsgHdr.RecursionDesired = true

	c := new(dns.Client)
	resp, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, answer := range resp.Answer {
		if rr, ok := answer.(*dns.A); ok {
			addresses = append(addresses, rr.A.String())
		} else if rr, ok := answer.(*dns.AAAA); ok {
			addresses = append(addresses, rr.AAAA.String())
		}
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf("failed! No valid A or AAAA record found for hostname: %s", hostname)
	}

	return addresses, nil
}
