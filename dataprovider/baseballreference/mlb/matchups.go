package mlb

import (
	"fmt"
	"log"
	"strings"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"
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

// MatchupOption defines a configuration option for MatchupRunner
type MatchupOption func(*MatchupRunner)

// WithMatchupTimeout sets the timeout duration for matchup runner
func WithMatchupTimeout(timeout time.Duration) MatchupOption {
	return func(mr *MatchupRunner) {
		mr.Timeout = timeout
	}
}

// WithMatchupDebug enables or disables debug mode for matchup runner
func WithMatchupDebug(debug bool) MatchupOption {
	return func(mr *MatchupRunner) {
		mr.Debug = debug
	}
}

// NewMatchupRunner creates a new MatchupRunner with the provided options
func NewMatchupRunner(options ...MatchupOption) *MatchupRunner {
	mr := &MatchupRunner{}

	// Apply all options
	for _, option := range options {
		option(mr)
	}

	return mr
}

// MatchupRunner specialized Runner for retrieving MLB matchup information.
type MatchupRunner struct {
	sportsreferenceutil.MatchupRunner
}

// GetMatchups retrieves MLB matchups for the specified date.
//
// Parameter:
//   - date: The date for which to retrieve matchups
//
// Returns a slice of MLB matchup data as interface{} values
func (matchupRunner *MatchupRunner) GetMatchups(date string) []interface{} {
	var matchups []interface{}
	timestamp, err := sportsreferenceutil.DateStrToTime(date)
	if err != nil {
		log.Fatalln(err)
	}
	month := timestamp.Format("1")
	day := timestamp.Format("2")
	year := timestamp.Format("2006")
	url, err := util.StrFormat(baseballreference.MatchupURL, "month", month, "year", year, "day", day)
	if err != nil {
		log.Fatalln(err)
	}
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	log.Println("Scraping Matchups: " + url)

	EventDate, err := sportsreferenceutil.EventDate(date)
	if err != nil {
		log.Fatalln(err)
	}
	doc, err := matchupRunner.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		log.Fatalln(err)
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
		if n == 3 {
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
				log.Fatalf("Game summary table #%d has %d table rows with unfamiliar record order!\n", idx+1, n)
			}
		} else if n == 2 {
			awayTeamSelection = teamSection[0]
			homeTeamSelection = teamSection[1]
			awayLocation = fmt.Sprintf(directTeamSelectorTemplate, 1)
			homeLocation = fmt.Sprintf(directTeamSelectorTemplate, 2)
		} else {
			log.Fatalf("Game summary table #%d has %d table rows. The expectation is either 2 or 3 rows!\n", idx+1, n)
		}

		// AwayTeam
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamNameSelector)
		awayName, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum((awayTeamSelection.Find(teamNameSelector).Text())), location, "AwayTeam")
		if err != nil {
			log.Fatalln(err)
		}
		if awayName == "National League" || awayName == "American League" {
			log.Printf("%s is a team associated with an all-star event. Skipping event date, %s, entirely", awayName, date)
			return false
		}
		matchup.AwayTeam = awayName

		// AwayTeamLink
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamLinkSelector)
		urlPath, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(teamLinkSelector).AttrOr("href", "")), location, "AwayTeamLink")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayTeamLink = baseballreference.URL + urlPath

		// AwayScore
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, teamScoreSelector)
		rawAwayScore, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(teamScoreSelector).Text()), location, "AwayScore")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayScore, err = util.TextToInt32(rawAwayScore)
		if err != nil {
			log.Printf("Cannot convert '%s' for rawAwayScore into Int\n", rawAwayScore)
			log.Fatalln(err)
		}

		// HomeTeam
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamNameSelector)
		homeName, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum((homeTeamSelection.Find(teamNameSelector).Text())), location, "HomeTeam")
		if err != nil {
			log.Fatalln(err)
		}
		if homeName == "National League" || homeName == "American League" {
			log.Printf("%s is a team associated with an all-star event. Skipping event date, %s, entirely", homeName, date)
			return false
		}
		matchup.HomeTeam = homeName

		// HomeTeamLink
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamLinkSelector)
		urlPath, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(homeTeamSelection.Find(teamLinkSelector).AttrOr("href", "")), location, "HomeTeamLink")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeTeamLink = baseballreference.URL + urlPath

		// HomeScore
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, homeLocation, teamScoreSelector)
		rawHomeScore, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(homeTeamSelection.Find(teamScoreSelector).Text()), location, "HomeScore")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeScore, err = util.TextToInt32(rawHomeScore)
		if err != nil {
			log.Printf("Cannot convert '%s' for rawHomeScore into Int\n", rawHomeScore)
			log.Fatalln(err)
		}

		// Loser
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation)
		rawLoser := util.CleanTextDatum(awayTeamSelection.AttrOr("class", ""))
		_, ok := sportsreferenceutil.LoserValueExists[rawLoser]
		if !ok {
			log.Fatalf("Unrecognized attribute value @ %s for Loser\n", location)
		}
		if rawLoser == "loser" {
			matchup.Loser = matchup.AwayTeam
		} else {
			matchup.Loser = matchup.HomeTeam
		}

		//BoxScoreLink
		location = fmt.Sprintf("%s %s %s %s", matchupsGameSummariesSelector, mlbGameSummarySelector, awayLocation, boxScoreLinkSelector)
		urlPath, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(awayTeamSelection.Find(boxScoreLinkSelector).AttrOr("href", "")), location, "BoxScoreLink")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.BoxScoreLink = baseballreference.URL + urlPath

		// EventID
		eventID, err := sportsreferenceutil.EventID(matchup.BoxScoreLink)
		if err != nil {
			log.Fatalln(err)
		}
		matchup.EventID = eventID

		matchups = append(matchups, matchup)
		return true
	})

	if len(matchups) == 0 {
		log.Printf("No relevant data scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		log.Printf("Scraping of %s Completed in %s\n", url, diff)
	}
	return matchups
}
