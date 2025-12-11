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
	doc, err := s.RetrieveDocument(URL, network.Headers{}, Selector)
	if err != nil {
		return "", err
	}
	log.Println("Document retrieved")
	return doc.Find(Selector).Text(), nil
}

func (s Scraper) Provider() sportscrape.Provider {
	return sportscrape.NBA
}
