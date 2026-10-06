package wnba

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

const (
	Selector = "script#__NEXT_DATA__"
	// UserAgent is shared by the chromedp-based box-score scrapers (via
	// NetworkHeaders) and the plain-HTTP matchup scraper (see
	// base_scraper_matchup.go) - wnba.com's api/schedule endpoint 403s
	// Go's default User-Agent, unlike the other JSON APIs this repo talks to.
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36"
)

var NetworkHeaders = network.Headers{
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
	"accept-language":           "en-US,en;q=0.8",
	"user-agent":                UserAgent,
	"sec-ch-ua-mobile":          "?0",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
}

// Scraper is the base for the chromedp/__NEXT_DATA__-based box-score scrapers.
type Scraper struct {
	scraper.BaseDocumentScraper
	// FetchAttempts is the number of times to retry fetching a page when the
	// fetch fails or the page's payload is rejected (see fetchDocWithRetry).
	// <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts.
	// <= 0 retries without a delay.
	FetchRetryBackoff time.Duration
}

func (s *Scraper) Provider() sportscrape.Provider {
	return sportscrape.WNBA
}

// Page loads intermittently hang until the timeout, and box score pages are
// sometimes served without player statistics, so fetchDocWithRetry retries a
// few times rather than treat one bad response as final. DefaultFetchAttempts
// is applied when a scraper doesn't set FetchAttempts explicitly (e.g. via the
// CLI flags).
const (
	DefaultFetchAttempts = 3
)

// fetchDocWithRetry fetches url, waiting for the __NEXT_DATA__ selector, and
// retries up to FetchAttempts times (sleeping FetchRetryBackoff between each)
// if the fetch fails or check rejects the document. A nil check accepts any
// document that was fetched. FetchAttempts <= 0 is treated as
// DefaultFetchAttempts; FetchRetryBackoff <= 0 retries without a delay.
func (s *Scraper) fetchDocWithRetry(url string, check func(*goquery.Document) error) (*goquery.Document, error) {
	attempts := s.FetchAttempts
	if attempts <= 0 {
		attempts = DefaultFetchAttempts
	}
	backoff := s.FetchRetryBackoff
	if backoff < 0 {
		backoff = 0
	}
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		doc, err := s.FetchDoc(url, Selector)
		switch {
		case err != nil:
			lastErr = err
		case check != nil:
			if lastErr = check(doc); lastErr == nil {
				return doc, nil
			}
		default:
			return doc, nil
		}
		if attempt < attempts {
			log.Printf("Attempt %d/%d fetching %s failed (%v); retrying in %s\n", attempt, attempts, url, lastErr, backoff)
			time.Sleep(backoff)
		}
	}
	return nil, lastErr
}
