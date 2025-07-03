package basketballreferencenba

import (
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"

	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// https://www.basketball-reference.com/boxscores/ Selectors for Matchups
	matchupsGameSummariesSelector = "#content >div.game_summaries"
	matchupsGameSummarySelector   = "div.game_summary"
	matchupsTeamsSelector         = "table.teams > tbody"
	matchupsAwayTeamSelector      = matchupsTeamsSelector + " > tr:nth-child(1)"
	matchupsAwayTeamNameSelector  = matchupsAwayTeamSelector + " > td:nth-child(1)"
	matchupsAwayTeamLinkSelector  = matchupsAwayTeamSelector + " > td:nth-child(1) > a"
	matchupsAwayTeamScoreSelector = matchupsAwayTeamSelector + " > td:nth-child(2)"
	matchupsHomeTeamSelector      = matchupsTeamsSelector + " > tr:nth-child(2)"
	matchupsHomeTeamNameSelector  = matchupsHomeTeamSelector + " > td:nth-child(1)"
	matchupsHomeTeamLinkSelector  = matchupsHomeTeamSelector + " > td:nth-child(1) > a"
	matchupsHomeTeamScoreSelector = matchupsHomeTeamSelector + " > td:nth-child(2)"
	matchupsLoserSelector         = matchupsAwayTeamSelector
	matchupsBoxScoreLinkSelector  = "p > a:nth-child(1)"

	matchupsAwayQuarterScoresSelector = "table:nth-child(2) > tbody > tr:nth-child(1) > td:nth-child(%d)"

	matchupsHomeQuarterScoresSelector = "table:nth-child(2) > tbody > tr:nth-child(2) > td:nth-child(%d)"

	matchupsQuarterHeadersSelector = "table:nth-child(2) > thead > tr > th:nth-child(%d)"
)

// MatchupOption defines a configuration option for MatchupScraper
type MatchupOption func(*MatchupScraper)

// WithMatchupTimeout sets the timeout duration for matchup runner
func WithMatchupDate(date string) MatchupOption {
	return func(mr *MatchupScraper) {
		mr.Date = date
	}
}

// WithMatchupTimeout sets the timeout duration for matchup runner
func WithMatchupTimeout(timeout time.Duration) MatchupOption {
	return func(mr *MatchupScraper) {
		mr.Timeout = timeout
	}
}

// WithMatchupDebug enables or disables debug mode for matchup runner
func WithMatchupDebug(debug bool) MatchupOption {
	return func(mr *MatchupScraper) {
		mr.Debug = debug
	}
}

// NewMatchupScraper creates a new MatchupScraper with the provided options
func NewMatchupScraper(options ...MatchupOption) *MatchupScraper {
	mr := &MatchupScraper{}

	// Apply all options
	for _, option := range options {
		option(mr)
	}
	mr.Init()

	return mr
}

// MatchupScraper specialized Runner for retrieving NBA matchup information.
type MatchupScraper struct {
	sportsreference.BaseScraper
	Date string
}

func (ms MatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.BasketballReference
}

func (ms MatchupScraper) Init() {
	ms.BaseScraper.Init()
	if ms.Date == "" {
		log.Fatalln("Date is a required argument")
	}
}

func (ms MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.BasketballReferenceNBAMatchup
}

// GetMatchups retrieves NBA matchups for the specified date.
func (ms *MatchupScraper) Scrape() sportscrape.MatchupOutput {
	var matchups []interface{}
	output := sportscrape.MatchupOutput{}
	timestamp, err := util.DateStrToTime(ms.Date)
	if err != nil {
		output.Error = err
		return output
	}
	month := timestamp.Format("1")
	day := timestamp.Format("2")
	year := timestamp.Format("2006")
	url, err := util.StrFormat(sportsreference.BasketballRefMatchupURL, "month", month, "year", year, "day", day)
	if err != nil {
		output.Error = err
		return output
	}
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	log.Println("Scraping Matchups: " + url)

	EventDate, err := sportsreference.EventDate(ms.Date)
	if err != nil {
		output.Error = err
		return output
	}
	doc, err := ms.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		output.Error = err
		return output
	}

	doc.Find(matchupsGameSummarySelector).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		var matchup model.NBAMatchup
		var location string
		matchup.PullTimestamp = PullTimestamp
		matchup.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		matchup.EventDate = EventDate
		matchup.EventDateParquet = util.TimeToDays(EventDate)
		// AwayTeam
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamNameSelector)

		matchup.AwayTeam, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum((s.Find(matchupsAwayTeamNameSelector).Text())), location, "AwayTeam")
		if err != nil {
			output.Error = err
			return false
		}

		// AwayTeamLink
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamLinkSelector)
		urlPath, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsAwayTeamLinkSelector).AttrOr("href", "")), location, "AwayTeamLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayTeamLink = sportsreference.BasketballRefURL + urlPath

		// AwayTeamID
		awayteamid, err := sportsreference.TeamID(urlPath)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayTeamID = awayteamid
		// AwayScore
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamScoreSelector)
		rawAwayScore, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsAwayTeamScoreSelector).Text()), location, "AwayScore")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayScore, err = util.TextToInt32(rawAwayScore)
		if err != nil {
			log.Printf("Cannot convert '%s' for rawAwayScore into Int\n", rawAwayScore)
			output.Error = err
			return false
		}

		// HomeTeam
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamNameSelector)
		matchup.HomeTeam, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamNameSelector).Text()), location, "HomeTeam")
		if err != nil {
			output.Error = err
			return false
		}
		// HomeTeamLink
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamLinkSelector)
		urlPath, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamLinkSelector).AttrOr("href", "")), location, "HomeTeamLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeTeamLink = sportsreference.BasketballRefURL + urlPath
		// HomeTeamID
		hometeamid, err := sportsreference.TeamID(urlPath)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeTeamID = hometeamid
		// HomeScore
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamScoreSelector)
		rawHomeScore, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamScoreSelector).Text()), location, "HomeScore")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeScore, err = util.TextToInt32(rawHomeScore)
		if err != nil {
			log.Printf("Conversion issue detected for rawHomeScore ('%s') to Int\n", rawHomeScore)
			output.Error = err
			return false
		}

		// Loser
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsLoserSelector)
		rawLoser := util.CleanTextDatum(s.Find(matchupsLoserSelector).AttrOr("class", ""))
		_, ok := sportsreference.LoserValueExists[rawLoser]
		if !ok {
			output.Error = fmt.Errorf("unrecognized attribute value @ %s for Loser", location)
			return false
		}
		if rawLoser == "loser" {
			matchup.Loser = matchup.AwayTeam
		} else {
			matchup.Loser = matchup.HomeTeam
		}

		//BoxScoreLink
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsBoxScoreLinkSelector)
		urlPath, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsBoxScoreLinkSelector).AttrOr("href", "")), location, "BoxScoreLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.BoxScoreLink = sportsreference.BasketballRefURL + urlPath

		// EventID
		eventID, err := sportsreference.EventID(matchup.BoxScoreLink)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.EventID = eventID

		// Quarter Headers check
		var selector string
		for _, position := range []int{2, 3, 4, 5} {
			selector = fmt.Sprintf(matchupsQuarterHeadersSelector, position)
			location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
			quarter := util.CleanTextDatum(s.Find(selector).Text())
			if quarter != fmt.Sprintf("%d", position-1) {
				log.Fatalf("Quarter %d not available @ %s\n", position-1, location)
			}
		}

		//AwayQ1Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 2)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ1TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ1TotalText, err = sportsreference.ReturnUnemptyField(AwayQ1TotalText, location, "AwayQ1Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayQ1Total, err = util.TextToInt32(AwayQ1TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ1TotalText to Int\n", AwayQ1TotalText)
			output.Error = err
			return false
		}

		//AwayQ2Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 3)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ2TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ2TotalText, err = sportsreference.ReturnUnemptyField(AwayQ2TotalText, location, "AwayQ2Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayQ2Total, err = util.TextToInt32(AwayQ2TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ2TotalText to Int\n", AwayQ2TotalText)
			output.Error = err
			return false
		}

		//AwayQ3Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 4)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ3TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ3TotalText, err = sportsreference.ReturnUnemptyField(AwayQ3TotalText, location, "AwayQ3Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayQ3Total, err = util.TextToInt32(AwayQ3TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ3TotalText to Int\n", AwayQ3TotalText)
			output.Error = err
			return false
		}

		//AwayQ4Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 5)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ4TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ4TotalText, err = sportsreference.ReturnUnemptyField(AwayQ4TotalText, location, "AwayQ4Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayQ4Total, err = util.TextToInt32(AwayQ4TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ4TotalText to Int\n", AwayQ4TotalText)
			output.Error = err
			return false
		}

		//HomeQ1Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 2)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ1TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ1TotalText, err = sportsreference.ReturnUnemptyField(HomeQ1TotalText, location, "HomeQ1Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeQ1Total, err = util.TextToInt32(HomeQ1TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ1TotalText to Int\n", HomeQ1TotalText)
			output.Error = err
			return false
		}

		//HomeQ2Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 3)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ2TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ2TotalText, err = sportsreference.ReturnUnemptyField(HomeQ2TotalText, location, "HomeQ2Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeQ2Total, err = util.TextToInt32(HomeQ2TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ2TotalText to Int\n", HomeQ2TotalText)
			output.Error = err
			return false
		}

		//HomeQ3Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 4)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ3TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ3TotalText, err = sportsreference.ReturnUnemptyField(HomeQ3TotalText, location, "HomeQ3Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeQ3Total, err = util.TextToInt32(HomeQ3TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ3TotalText to Int\n", HomeQ3TotalText)
			output.Error = err
			return false
		}

		//HomeQ4Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 5)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ4TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ4TotalText, err = sportsreference.ReturnUnemptyField(HomeQ4TotalText, location, "HomeQ4Total")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeQ4Total, err = util.TextToInt32(HomeQ4TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ4TotalText to Int\n", HomeQ4TotalText)
			output.Error = err
			return false
		}

		matchups = append(matchups, matchup)
		return true
	})

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Completed in %s\n", url, diff)
	output.Output = matchups
	return output
}
