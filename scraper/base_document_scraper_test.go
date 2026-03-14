//go:build unit

package scraper

import (
	"testing"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape/util/request"
	"github.com/stretchr/testify/assert"
)

// TestBaseDocumentScraperStructure verifies that BaseDocumentScraper stores
// its configuration fields correctly.
func TestBaseDocumentScraperStructure(t *testing.T) {
	testCases := []struct {
		name    string
		timeout time.Duration
		debug   bool
	}{
		{"Zero values", 0, false},
		{"WithTimeout", 5 * time.Second, false},
		{"WithDebug", 0, true},
		{"WithTimeoutAndDebug", 10 * time.Second, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := BaseDocumentScraper{
				Timeout: tc.timeout,
				Debug:   tc.debug,
			}
			assert.Equal(t, tc.timeout, s.Timeout)
			assert.Equal(t, tc.debug, s.Debug)
		})
	}
}

// TestBaseDocumentScraperInitNoOp verifies that Init is a no-op when
// DocumentRetriever is already set.
func TestBaseDocumentScraperInitNoOp(t *testing.T) {
	existing := &request.DocumentRetrieverV2{}
	s := &BaseDocumentScraper{
		DocumentRetriever: existing,
	}
	s.Init()
	assert.Equal(t, existing, s.DocumentRetriever, "Init should not replace an existing DocumentRetriever")
}

// TestBaseDocumentScraperFetchDocNilRetriever verifies that FetchDoc returns an
// error when DocumentRetriever has not been initialised.
func TestBaseDocumentScraperFetchDocNilRetriever(t *testing.T) {
	s := &BaseDocumentScraper{}
	doc, err := s.FetchDoc("https://example.com", "body")
	assert.Error(t, err)
	assert.Nil(t, doc)
	assert.Contains(t, err.Error(), "DocumentRetriever is required")
}

// TestBaseDocumentScraperNetworkHeaders verifies that NetworkHeaders is stored
// on the struct.
func TestBaseDocumentScraperNetworkHeaders(t *testing.T) {
	headers := network.Headers{"User-Agent": "test-agent"}
	s := &BaseDocumentScraper{
		NetworkHeaders: headers,
	}
	assert.Equal(t, headers, s.NetworkHeaders)
}

// TestBaseDocumentScraperCloseNilRetriever verifies that Close is safe to call
// when DocumentRetriever is nil.
func TestBaseDocumentScraperCloseNilRetriever(t *testing.T) {
	s := &BaseDocumentScraper{}
	assert.NotPanics(t, func() { s.Close() })
	assert.Nil(t, s.DocumentRetriever)
}

// TestBaseDocumentScraperCloseNonNilRetriever verifies that Close sets
// DocumentRetriever to nil after tearing down the session.
func TestBaseDocumentScraperCloseNonNilRetriever(t *testing.T) {
	s := &BaseDocumentScraper{
		DocumentRetriever: &request.DocumentRetrieverV2{},
	}
	s.Close()
	assert.Nil(t, s.DocumentRetriever, "Close should set DocumentRetriever to nil")
}
