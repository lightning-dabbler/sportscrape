package baseballreferencemlb

import (
	"fmt"
	"log"
	"strings"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/scraper"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"

	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// https://www.baseball-reference.com/boxes/ Selectors for Matchups
	matchupsGameSummariesSelector = "#content >div.game_summaries"
	mlbGameSummarySelector        = "div.game_summary"
	matchupTeamsSelector          = "table.teams > tbody tr"
	directTeamSelectorTemplate    = " table.teams > tbody > tr:nth-child(%d)"
	teamNameSelector              = "td:nth-child(1)"
	teamLinkSelector              = "td:nth-child(1) > a"
	teamScoreSelector             = "td:nth-child(2)"
	boxScoreLinkSelector          = "td.right.gamelink > a"
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

// MatchupScraper specialized Runner for retrieving MLB matchup information.
type MatchupScraper struct {
	scraper.BaseScraper
	Date string
}

func (ms MatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.BaseballReference
}

func (ms MatchupScraper) Init() {
	ms.BaseScraper.Init()
	if ms.Date == "" {
		log.Fatalln("Date is a required argument")
	}
}

func (ms MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballReferenceMLBMatchup
}

// Scrape retrieves MLB matchups for the specified date.
func (ms *MatchupScraper) Scrape() sportscrape.MatchupOutput {
	var matchups []interface{}
	output := sportscrape.MatchupOutput{}
	var skips int
	timestamp, err := util.DateStrToTime(ms.Date)
	if err != nil {
		output.Error = err
		return output
	}
	month := timestamp.Format("1")
	day := timestamp.Format("2")
	year := timestamp.Format("2006")
	url, err := util.StrFormat(sportsreference.BaseballRefMatchupURL, "month", month, "year", year, "day", day)
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

	doc.Find(mlbGameSummarySelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		var matchup model.MLBMatchup
		var location, homeLocation, awayLocation string
		matchup.PullTimestamp = PullTimestamp
		matchup.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		matchup.EventDate = EventDate
		matchup.EventDateParquet = util.TimeToDays(EventDate)
		// Teams
		var teamSection []*goquery.Selection
		s.Find(matchupTeamsSelector).Each(func(_ int, s *goquery.Selection) {
			teamSection = append(teamSection, s)
		})
		var awayTeamSelection, homeTeamSelection *goquery.Selection
		n := len(teamSection)
		switch n {
		case 3:
			if strings.HasPrefix(strings.ToLower(teamSection[0].Find("td").Text()), "game") {
				// playoff event
				matchup.PlayoffMatch = true
				awayTeamSelection = teamSection[1]
				homeTeamSelection = teamSection[2]
				awayLocation = fmt.Sprintf(directTeamSelectorTemplate, 2)
				homeLocation = fmt.Sprintf(directTeamSelectorTemplate, 3)
			} else if strings.HasPrefix(strings.ToLower(teamSection[2].Find("td").Text()), "completed on") {
				// The game completed at a later date
				awayTeamSelection = teamSection[0]
				homeTeamSelection = teamSection[1]
				awayLocation = fmt.Sprintf(directTeamSelectorTemplate, 1)
				homeLocation = fmt.Sprintf(directTeamSelectorTemplate, 2)
			} else {
				output.Error = fmt.Errorf("game summary table #%d has %d table rows with unfamiliar record order", idx+1, n)
				return false
			}
		case 2:
			awayTeamSelection = teamSection[0]
			homeTeamSelection = teamSection[1]
			awayLocation = fmt.Sprintf(directTeamSelectorTemplate, 1)
			homeLocation = fmt.Sprintf(directTeamSelectorTemplate, 2)
		default:
			output.Error = fmt.Errorf("game summary table #%d has %d table rows. The expectation is either 2 or 3 rows", idx+1, n)
			return false
		}
		// AwayTeam
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamNameSelector)
		awayName, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum((awayTeamSelection.Find(teamNameSelector).Text())), location, "AwayTeam")
		if err != nil {
			output.Error = err
			return false
		}
		if awayName == "National League" || awayName == "American League" {
			log.Printf("%s is a team associated with an all-star event. Skipping event date, %s, entirely\n", awayName, ms.Date)
			skips += 1
			return false
		}
		matchup.AwayTeam = awayName

		// AwayTeamLink
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamLinkSelector)
		urlPath, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(teamLinkSelector).AttrOr("href", "")), location, "AwayTeamLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayTeamLink = sportsreference.BaseballRefURL + urlPath
		// AwayTeamID
		awayteamid, err := sportsreference.TeamID(urlPath)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.AwayTeamID = awayteamid

		// AwayScore
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamScoreSelector)
		rawAwayScore, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(teamScoreSelector).Text()), location, "AwayScore")
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
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamNameSelector)
		homeName, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum((homeTeamSelection.Find(teamNameSelector).Text())), location, "HomeTeam")
		if err != nil {
			output.Error = err
			return false
		}
		if homeName == "National League" || homeName == "American League" {
			log.Printf("%s is a team associated with an all-star event. Skipping event date, %s, entirely\n", homeName, ms.Date)
			skips += 1
			return false
		}
		matchup.HomeTeam = homeName

		// HomeTeamLink
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamLinkSelector)
		urlPath, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum(homeTeamSelection.Find(teamLinkSelector).AttrOr("href", "")), location, "HomeTeamLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeTeamLink = sportsreference.BaseballRefURL + urlPath
		// HomeTeamID
		hometeamid, err := sportsreference.TeamID(urlPath)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeTeamID = hometeamid

		// HomeScore
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamScoreSelector)
		rawHomeScore, err := sportsreference.ReturnUnemptyField(util.CleanTextDatum(homeTeamSelection.Find(teamScoreSelector).Text()), location, "HomeScore")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.HomeScore, err = util.TextToInt32(rawHomeScore)
		if err != nil {
			log.Printf("Cannot convert '%s' for rawHomeScore into Int\n", rawHomeScore)
			output.Error = err
			return false
		}

		// Loser
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation)
		rawLoser := util.CleanTextDatum(awayTeamSelection.AttrOr("class", ""))
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
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, boxScoreLinkSelector)
		urlPath, err = sportsreference.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(boxScoreLinkSelector).AttrOr("href", "")), location, "BoxScoreLink")
		if err != nil {
			output.Error = err
			return false
		}
		matchup.BoxScoreLink = sportsreference.BaseballRefURL + urlPath

		// EventID
		eventID, err := sportsreference.EventID(matchup.BoxScoreLink)
		if err != nil {
			output.Error = err
			return false
		}
		matchup.EventID = eventID

		matchups = append(matchups, matchup)
		return true
	})

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Completed in %s\n", url, diff)
	output.Output = matchups
	return output
}
