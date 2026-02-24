package nba

import (
	"log"

	"github.com/chromedp/cdproto/network"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

const (
	Selector = "script#__NEXT_DATA__"
)

type Scraper struct {
	scraper.BaseScraper
}

func (s Scraper) FetchDoc(URL string) (string, error) {
	doc, err := s.RetrieveDocument(URL, network.Headers{
		"user-agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
		"sec-ch-ua":        `"Chromium";v="141", "Google Chrome";v="141", "Not.A/Brand";v="99"`,
		"sec-ch-ua-mobile": "?0",
		"accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
		"sec-fetch-dest":   "document",
		"sec-fetch-mode":   "navigate",
		"sec-fetch-site":   "same-origin",
		"sec-fetch-user":   "?1",
		"referer":          "https://www.nba.com/",
	}, Selector)
	if err != nil {
		return "", err
	}
	log.Println("Document retrieved")
	return doc.Find(Selector).Text(), nil
}

func (s Scraper) Provider() sportscrape.Provider {
	return sportscrape.NBA
}
