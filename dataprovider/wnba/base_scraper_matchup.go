package wnba

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/scraper"
	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	BaseURL = "https://www.wnba.com/api/schedule"
)

// BaseMatchupScraper wraps scraper.BaseJsonScraper[jsonresponse.MatchupJSON]
// (plain HTTP, no browser) and exposes the same Date-based interface as
// NBA's BaseMatchupScraper, plus an optional EndDate for range queries.
// Internally it derives the season(s) to fetch from Date/EndDate's year(s)
// and filters the (always full-season) response down to the requested range.
type BaseMatchupScraper struct {
	scraper.BaseJsonScraper[jsonresponse.MatchupJSON]
	Date    string // YYYY-MM-DD, required
	EndDate string // YYYY-MM-DD, optional; empty = same as Date (single day)
}

func (bms *BaseMatchupScraper) Init() {
	if bms.Date == "" {
		log.Fatalln("Date is a required argument")
	}
	startT, err := util.DateStrToTime(bms.Date)
	if err != nil {
		log.Fatalln(err)
	}
	if bms.EndDate != "" {
		endT, err := util.DateStrToTime(bms.EndDate)
		if err != nil {
			log.Fatalln(err)
		}
		if endT.Before(startT) {
			log.Fatalln("EndDate must not be before Date")
		}
	}
	// scraper.BaseJsonScraper[T].Init() is a no-op; called for interface parity
	bms.BaseJsonScraper.Init()
}

// Close is a no-op: scraper.BaseJsonScraper has no Close(), and there is no
// browser session for this scraper to tear down.
func (bms *BaseMatchupScraper) Close() {}

func (bms *BaseMatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.WNBA
}

// dateRange returns the inclusive [start, end] range (end == start if EndDate is unset).
func (bms *BaseMatchupScraper) dateRange() (time.Time, time.Time, error) {
	start, err := util.DateStrToTime(bms.Date)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	if bms.EndDate == "" {
		return start, start, nil
	}
	end, err := util.DateStrToTime(bms.EndDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return start, end, nil
}

// seasons returns the distinct calendar years (as strings) spanned by the
// date range, in ascending order. Almost always a single year; handled
// defensively in case a range crosses a calendar year boundary.
func (bms *BaseMatchupScraper) seasons() ([]string, error) {
	start, end, err := bms.dateRange()
	if err != nil {
		return nil, err
	}
	var seasons []string
	for y := start.Year(); y <= end.Year(); y++ {
		seasons = append(seasons, strconv.Itoa(y))
	}
	return seasons, nil
}

// retrieveModel fetches and parses the schedule JSON at url. It bypasses
// scraper.BaseJsonScraper.RetrieveModel (which sends no headers at all, via
// util/request.Get) because wnba.com's api/schedule endpoint 403s requests
// without a browser-like User-Agent; HydrateModel is still reused for the
// actual JSON unmarshaling.
func (bms *BaseMatchupScraper) retrieveModel(url string) (*jsonresponse.MatchupJSON, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Error at %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to '%s' received a %s status", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bms.HydrateModel(body)
}

// URL builds https://www.wnba.com/api/schedule?season={season}&regionId=1
func (bms *BaseMatchupScraper) URL(season string) (string, error) {
	URL, err := url.Parse(BaseURL)
	if err != nil {
		return "", err
	}
	queryValues := URL.Query()
	queryValues.Add("season", season)
	queryValues.Add("regionId", "1")
	URL.RawQuery = queryValues.Encode()
	return URL.String(), nil
}
