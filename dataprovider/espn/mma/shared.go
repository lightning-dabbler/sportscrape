package mma

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
)

const espnFittMarker = "window['__espnfitt__']="

// ESPN intermittently serves a bot-check interstitial in place of the real
// page; the interstitial still satisfies the "html" ready selector, so
// fetchESPNFittPayload retries a few times rather than treat one empty
// response as final. These are the defaults applied when a scraper doesn't
// set FetchAttempts/FetchRetryBackoff explicitly (e.g. via the CLI flags).
const (
	DefaultFetchAttempts     = 3
	DefaultFetchRetryBackoff = 3 * time.Second
)

// extractESPNFittPayload returns the raw JSON payload embedded in the page's
// window['__espnfitt__'] script tag, if present.
func extractESPNFittPayload(doc *goquery.Document) ([]byte, bool) {
	var payload []byte
	doc.Find("script").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		text := s.Text()
		if strings.Contains(text, espnFittMarker) {
			parts := strings.SplitAfter(text, espnFittMarker)
			payload = []byte(parts[1][0 : len(parts[1])-1])
			return false
		}
		return true
	})
	return payload, payload != nil
}

// fetchESPNFittPayload fetches url via fetchDoc, waiting for selector, and
// retries up to attempts times (sleeping backoff between each) if the
// espnfitt payload isn't found in the response. attempts <= 0 is treated as
// DefaultFetchAttempts; backoff <= 0 is treated as DefaultFetchRetryBackoff.
func fetchESPNFittPayload(fetchDoc func(url, selector string) (*goquery.Document, error), url, selector string, attempts int, backoff time.Duration) ([]byte, error) {
	if attempts <= 0 {
		attempts = DefaultFetchAttempts
	}
	if backoff <= 0 {
		backoff = DefaultFetchRetryBackoff
	}
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		doc, err := fetchDoc(url, selector)
		switch {
		case err != nil:
			lastErr = err
		default:
			if payload, ok := extractESPNFittPayload(doc); ok {
				return payload, nil
			}
			lastErr = fmt.Errorf("espnfitt payload not found in response from %s", url)
		}
		if attempt < attempts {
			log.Printf("Attempt %d/%d fetching %s failed (%v); retrying in %s\n", attempt, attempts, url, lastErr, backoff)
			time.Sleep(backoff)
		}
	}
	return nil, lastErr
}

var NetworkHeaders network.Headers = network.Headers(map[string]any{
	"authority":       "www.espn.com",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
	"accept-language": "en-US;q=0.8",
	// "cookie":                    "is_live=true; sr_note_box_countdown=57",
	// "if-modified-since":         "Tue, 08 Nov 2022 01:08:31 GMT",
	// "sec-fetch-dest":            "document",
	// "sec-fetch-mode":            "navigate",
	// "sec-fetch-site":            "none",
	// "sec-fetch-user":            "?1",
	"sec-ch-ua-mobile":          "?0",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
})
