package request

import (
	"fmt"
	"log"
	"net/http"
)

// Get performs a GET request using the url it receives
// Returns an http response
func Get(url string) *http.Response {
	fmt.Printf("Fetching from %s\n", url)
	resp, httpGetErr := http.Get(url)
	if httpGetErr != nil {
		log.Printf("HTTP Error at %s\n", url)
		log.Fatalln(httpGetErr)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Request to '%s' %s\n", url, resp.Status)
	}
	return resp
}
