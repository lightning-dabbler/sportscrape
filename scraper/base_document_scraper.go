package scraper

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

// BaseDocumentScraper provides base functionality for scraping web documents
// using a persistent headless browser session (DocumentRetrieverV2).
// Call Init() before use and Close() when done to release browser resources.
type BaseDocumentScraper struct {
	DocumentRetriever *request.DocumentRetrieverV2
	Timeout           time.Duration
	Debug             bool
	NetworkHeaders    network.Headers
}

// Init initializes the DocumentRetriever using the scraper's Timeout, Debug,
// and NetworkHeaders fields. If DocumentRetriever is already set, Init is a
// no-op. Calls log.Fatalln if Timeout is zero and no DocumentRetriever is set.
func (s *BaseDocumentScraper) Init() {
	if s.DocumentRetriever == nil {
		if s.Timeout == 0 {
			log.Fatalln("Timeout needs to be > 0")
		}

		documentretriever, err := request.NewDocumentRetrieverV2(
			s.NetworkHeaders,
			request.WithDebugV2(s.Debug),
			request.WithTimeoutV2(s.Timeout),
		)
		if err != nil {
			log.Fatalln(err)
		}
		s.DocumentRetriever = documentretriever
	}
}
// FetchDoc retrieves and parses the HTML document at URL, waiting for the
// element matched by selector to be ready before extracting its outer HTML.
//
// Parameters:
//   - URL: The web page URL to retrieve
//   - selector: CSS selector of the element to wait for and extract
//
// Returns an error if DocumentRetriever is nil or if retrieval or parsing fails.
func (s *BaseDocumentScraper) FetchDoc(URL string, selector string) (*goquery.Document, error) {
	if s.DocumentRetriever == nil {
		return nil, fmt.Errorf("DocumentRetriever is required")
	}
	doc, err := s.DocumentRetriever.RetrieveDocument(URL, selector)
	if err != nil {
		return nil, err
	}
	log.Println("Document retrieved")
	return doc, nil
}

// Close tears down the underlying browser session and sets DocumentRetriever
// to nil. Safe to call when DocumentRetriever is nil.
func (s *BaseDocumentScraper) Close() {
	if s.DocumentRetriever != nil {
		s.DocumentRetriever.Close()
		s.DocumentRetriever = nil
	}
}
