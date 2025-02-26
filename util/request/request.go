package request

import (
	"fmt"
	"net/http"
)

// Get performs a GET request using the url it receives
// Returns an http response
func Get(url string) (*http.Response, error) {
	fmt.Printf("Fetching from %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request to '%s' received a %s status", url, resp.Status)
	}
	return resp, nil
}
