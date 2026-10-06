package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
)

type BaseEventDataScraper struct {
	Scraper
	Period       Period
	FeedType     FeedType
	BoxScoreType BoxScoreType
	// FetchAttempts is the number of times to retry fetching a box score page
	// when the fetch fails or the payload is missing player statistics.
	// <= 0 falls back to DefaultFetchAttempts.
	FetchAttempts int
	// FetchRetryBackoff is the delay between retry attempts.
	// <= 0 retries without a delay.
	FetchRetryBackoff time.Duration
}

func (beds *BaseEventDataScraper) Init() {
	if beds.FeedType.Undefined() {
		log.Fatalln("FeedType is a required argument")
	}

	switch beds.FeedType {
	case BoxScore:
		if beds.BoxScoreType.Undefined() {
			log.Fatalln("BoxScoreType is a required argument when FeedType is BoxScore")
		}
		switch beds.BoxScoreType {
		case Traditional, Advanced, Misc, Scoring, Usage, FourFactors:
			if beds.Period.Undefined() {
				log.Printf("Warning: Period is unset for BoxScore FeedType (%s)... defaulting to %s\n", beds.BoxScoreType.Type(), Full.Period())
				beds.Period = Full
			}
		}
	case PlayByPlay:
		if !beds.BoxScoreType.Undefined() {
			log.Println("Warning: BoxScoreType argument will be ignored for PlayByPlay FeedType")
		}
		if beds.Period != Full {
			log.Println("Setting Period to Full for PlayByPlay FeedType")
			beds.Period = Full
		}
	}
	beds.Scraper.Init()
}

func (beds BaseEventDataScraper) ConstructContext(matchup model.Matchup) sportscrape.EventDataContext {
	return sportscrape.EventDataContext{
		AwayTeam:  matchup.AwayTeam,
		AwayID:    matchup.AwayTeamID,
		HomeTeam:  matchup.HomeTeam,
		HomeID:    matchup.HomeTeamID,
		EventTime: matchup.EventTime,
		EventID:   matchup.EventID,
	}
}

func (beds BaseEventDataScraper) URL(share_url string) (string, error) {
	if share_url == "" {
		return "", fmt.Errorf("share_url should not be empty")
	}
	URL, err := url.Parse(share_url)
	if err != nil {
		return "", err
	}
	var urlstr string
	switch beds.FeedType {
	case BoxScore:
		joinedURL := URL.JoinPath(BoxScore.Type())
		queryValues := joinedURL.Query()
		switch beds.BoxScoreType {
		case Traditional, Advanced, Misc, Scoring, Usage, FourFactors:
			queryValues.Add("period", beds.Period.Period())
		}
		if beds.BoxScoreType != Live {
			queryValues.Add("type", beds.BoxScoreType.Type())
		}
		joinedURL.RawQuery = queryValues.Encode()
		urlstr = joinedURL.String()

	case PlayByPlay:
		joinedURL := URL.JoinPath(PlayByPlay.Type())
		queryValues := joinedURL.Query()
		queryValues.Add("period", beds.Period.Period())
		joinedURL.RawQuery = queryValues.Encode()
		urlstr = joinedURL.String()
	}

	if urlstr == "" {
		return "", fmt.Errorf("result urlstr should not be empty")
	}
	return urlstr, nil
}

func (beds BaseEventDataScraper) PeriodBasedBoxScoreDataAvailable(period int32, gameStatus int32) bool {
	// Game is Final
	if gameStatus == int32(3) {
		switch beds.Period {
		case AllOT:
			if period > 4 {
				return true
			}
		default:
			return true
		}
	}
	return false
}

func (beds BaseEventDataScraper) NonPeriodBasedBoxScoreDataAvailable(gameStatus int32) bool {
	// Game is Final
	if gameStatus == int32(3) {
		return true
	}
	return false
}

func (beds BaseEventDataScraper) LiveBoxScoreDataAvailable(gameStatus int32) bool {
	// Game is ongoing
	if gameStatus == int32(2) {
		return true
	}
	return false
}

// nba.com and wnba.com intermittently server-render box score pages whose
// player objects carry only names and IDs, with no statistics. The payload
// still decodes, so fetchBoxScorePayload retries a few times rather than emit
// rows of zero values. DefaultFetchAttempts is applied when a scraper doesn't
// set FetchAttempts explicitly (e.g. via the CLI flags).
const (
	DefaultFetchAttempts = 3
)

// fetchBoxScorePayload fetches url, waiting for the __NEXT_DATA__ selector, and
// retries up to FetchAttempts times (sleeping FetchRetryBackoff between each)
// if the fetch fails or the players in the payload are missing statsKey.
// FetchAttempts <= 0 is treated as DefaultFetchAttempts; FetchRetryBackoff <= 0
// retries without a delay.
func (beds *BaseEventDataScraper) fetchBoxScorePayload(url, statsKey string) (string, error) {
	attempts := beds.FetchAttempts
	if attempts <= 0 {
		attempts = DefaultFetchAttempts
	}
	backoff := beds.FetchRetryBackoff
	if backoff < 0 {
		backoff = 0
	}
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		doc, err := beds.FetchDoc(url, Selector)
		switch {
		case err != nil:
			lastErr = err
		default:
			jsonstr := doc.Find(Selector).Text()
			if !boxScorePlayerStatsMissing(jsonstr, statsKey) {
				return jsonstr, nil
			}
			lastErr = fmt.Errorf("box score payload from %s is missing player %s", url, statsKey)
		}
		if attempt < attempts {
			log.Printf("Attempt %d/%d fetching %s failed (%v); retrying in %s\n", attempt, attempts, url, lastErr, backoff)
			time.Sleep(backoff)
		}
	}
	return "", lastErr
}

// boxScorePlayerStatsMissing reports whether a live or final game's payload
// lists players but none of them carry statsKey. Unparseable payloads report
// false so the caller's own decoding surfaces the error.
func boxScorePlayerStatsMissing(jsonstr, statsKey string) bool {
	var payload struct {
		Props struct {
			PageProps struct {
				Game struct {
					GameStatus int32 `json:"gameStatus"`
					HomeTeam   struct {
						Players []map[string]json.RawMessage `json:"players"`
					} `json:"homeTeam"`
					AwayTeam struct {
						Players []map[string]json.RawMessage `json:"players"`
					} `json:"awayTeam"`
				} `json:"game"`
			} `json:"pageProps"`
		} `json:"props"`
	}
	if err := json.Unmarshal([]byte(jsonstr), &payload); err != nil {
		return false
	}
	game := payload.Props.PageProps.Game
	if game.GameStatus != int32(2) && game.GameStatus != int32(3) {
		return false
	}
	players := append(game.HomeTeam.Players, game.AwayTeam.Players...)
	if len(players) == 0 {
		return false
	}
	for _, player := range players {
		if _, ok := player[statsKey]; ok {
			return false
		}
	}
	return true
}
