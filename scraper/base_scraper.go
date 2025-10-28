package scraper

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

// BaseScraper provides base functionality for retrieving documents from web pages.
// It encapsulates configuration for document retrieval operations.
type BaseScraper struct {
	// Timeout is the context timeout for Chrome operations
	Timeout time.Duration
	// Debug enables verbose logging of Chrome operations when true
	Debug bool
}

func (s BaseScraper) Init() {
	if s.Timeout == 0 {
		log.Fatalln("Timeout needs to be > 0")
	}
}

// RetrieveDocument fetches and parses a web document from the specified URL.
// It uses a DocumentRetriever to load the page and parse it into a goquery Document.
//
// Parameters:
//   - url: The web page URL to retrieve
//   - networkHeaders: HTTP headers to include with the request
//   - waitReadySelector: CSS selector of element that indicates page is fully loaded
//
// Returns:
//   - *goquery.Document: The parsed document
//   - error: Any error encountered during fetching or parsing
func (s *BaseScraper) RetrieveDocument(url string, networkHeaders network.Headers, waitReadySelector string) (*goquery.Document, error) {
	var retrieveOptions []request.RetrieverOption
	// Only add the timeout option if it's non-zero
	if s.Timeout != 0 {
		retrieveOptions = append(retrieveOptions, request.WithTimeout(s.Timeout))
	}

	// Add debug option
	if s.Debug {
		retrieveOptions = append(retrieveOptions, request.WithDebug(s.Debug))
	}

	// Create the document retriever with the collected options
	dr := request.NewDocumentRetriever(retrieveOptions...)
	return dr.RetrieveDocument(url, networkHeaders, waitReadySelector)
}
