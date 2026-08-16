package wnba

import (
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
}

func (s *Scraper) Provider() sportscrape.Provider {
	return sportscrape.WNBA
}
