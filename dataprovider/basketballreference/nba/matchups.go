package nba

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"

	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
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

// GetMatchups accepts a date string in the form of 2024-01-25
// It fetches NBA matchup details that've occured on the date of interest
// Returns an array of model.NBAMatchup in the form of interface{}
func GetMatchups(date string) []interface{} {
	var matchups []interface{}
	timestamp, err := sportsreferenceutil.DateStrToTime(date)
	if err != nil {
		log.Fatalln(err)
	}
	month := timestamp.Format("1")
	day := timestamp.Format("2")
	year := timestamp.Format("2006")
	url, err := util.StrFormat(basketballreference.MatchupURL, "month", month, "year", year, "day", day)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	fmt.Println("Scraping Matchups: " + url)
	var outer string
	PullTimestamp := time.Now().UTC()
	EventDate, err := sportsreferenceutil.EventDate(date)
	if err != nil {
		log.Fatalln(err)
	}
	start := time.Now().UTC()
	if err := chromedp.Run(ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"authority":                 "www.basketball-reference.com",
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"accept-language":           "en-US;q=0.8",
			"cookie":                    "is_live=true; sr_note_box_countdown=57",
			"if-modified-since":         "Tue, 08 Nov 2022 01:08:31 GMT",
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"sec-gpc":                   "1",
			"upgrade-insecure-requests": "1",
			"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
		})),
		chromedp.Navigate(url),
		chromedp.WaitReady(matchupsGameSummariesSelector),
		chromedp.OuterHTML(matchupsGameSummariesSelector, &outer, chromedp.ByQuery),
	); err != nil && strings.Trim(err.Error(), " ") != "context deadline exceeded" {
		fmt.Println("ERROR:")
		fmt.Printf("GetMatchups %s\n", url)
		log.Fatalln(err)
	}
	myReader := strings.NewReader(outer)
	doc, err := goquery.NewDocumentFromReader(myReader)
	if err != nil {
		log.Fatalln(err)
	}

	doc.Find(matchupsGameSummarySelector).Each(func(_ int, s *goquery.Selection) {
		var matchup model.NBAMatchup
		var location string
		matchup.PullTimestamp = PullTimestamp
		matchup.EventDate = EventDate
		// AwayTeam
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamNameSelector)

		matchup.AwayTeam, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum((s.Find(matchupsAwayTeamNameSelector).Text())), location, "AwayTeam")
		if err != nil {
			log.Fatalln(err)
		}
		// AwayTeamLink
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamLinkSelector)
		urlPath, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsAwayTeamLinkSelector).AttrOr("href", "")), location, "AwayTeamLink")
		if err != nil {
			log.Fatalln(err)
		}

		matchup.AwayTeamLink = basketballreference.URL + urlPath
		// AwayScore
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsAwayTeamScoreSelector)
		rawAwayScore, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsAwayTeamScoreSelector).Text()), location, "AwayScore")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayScore, err = util.TextToInt(rawAwayScore)
		if err != nil {
			log.Printf("Cannot convert '%s' for rawAwayScore into Int\n", rawAwayScore)
			log.Fatalln(err)
		}

		// HomeTeam
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamNameSelector)
		matchup.HomeTeam, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamNameSelector).Text()), location, "HomeTeam")
		if err != nil {
			log.Fatalln(err)
		}
		// HomeTeamLink
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamLinkSelector)
		urlPath, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamLinkSelector).AttrOr("href", "")), location, "HomeTeamLink")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeTeamLink = basketballreference.URL + urlPath
		// HomeScore
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsHomeTeamScoreSelector)
		rawHomeScore, err := sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsHomeTeamScoreSelector).Text()), location, "HomeScore")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeScore, err = util.TextToInt(rawHomeScore)
		if err != nil {
			log.Printf("Conversion issue detected for rawHomeScore ('%s') to Int\n", rawHomeScore)
			log.Fatalln(err)
		}

		// Loser
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsLoserSelector)
		rawLoser := util.CleanTextDatum(s.Find(matchupsLoserSelector).AttrOr("class", ""))
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
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, matchupsBoxScoreLinkSelector)
		urlPath, err = sportsreferenceutil.ReturnUnemptyField(util.CleanTextDatum(s.Find(matchupsBoxScoreLinkSelector).AttrOr("href", "")), location, "BoxScoreLink")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.BoxScoreLink = basketballreference.URL + urlPath

		// EventID
		splitLink := strings.Split(matchup.BoxScoreLink, "/")
		matchup.EventID = strings.Split(splitLink[len(splitLink)-1], ".")[0]

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
		AwayQ1TotalText, err = sportsreferenceutil.ReturnUnemptyField(AwayQ1TotalText, location, "AwayQ1Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayQ1Total, err = util.TextToInt(AwayQ1TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ1TotalText to Int\n", AwayQ1TotalText)
			log.Fatalln(err)
		}

		//AwayQ2Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 3)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ2TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ2TotalText, err = sportsreferenceutil.ReturnUnemptyField(AwayQ2TotalText, location, "AwayQ2Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayQ2Total, err = util.TextToInt(AwayQ2TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ2TotalText to Int\n", AwayQ2TotalText)
			log.Fatalln(err)
		}

		//AwayQ3Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 4)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ3TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ3TotalText, err = sportsreferenceutil.ReturnUnemptyField(AwayQ3TotalText, location, "AwayQ3Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayQ3Total, err = util.TextToInt(AwayQ3TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ3TotalText to Int\n", AwayQ3TotalText)
			log.Fatalln(err)
		}

		//AwayQ4Total
		selector = fmt.Sprintf(matchupsAwayQuarterScoresSelector, 5)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		AwayQ4TotalText := util.CleanTextDatum(s.Find(selector).Text())
		AwayQ4TotalText, err = sportsreferenceutil.ReturnUnemptyField(AwayQ4TotalText, location, "AwayQ4Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.AwayQ4Total, err = util.TextToInt(AwayQ4TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for AwayQ4TotalText to Int\n", AwayQ4TotalText)
			log.Fatalln(err)
		}

		//HomeQ1Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 2)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ1TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ1TotalText, err = sportsreferenceutil.ReturnUnemptyField(HomeQ1TotalText, location, "HomeQ1Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeQ1Total, err = util.TextToInt(HomeQ1TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ1TotalText to Int\n", HomeQ1TotalText)
			log.Fatalln(err)
		}

		//HomeQ2Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 3)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ2TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ2TotalText, err = sportsreferenceutil.ReturnUnemptyField(HomeQ2TotalText, location, "HomeQ2Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeQ2Total, err = util.TextToInt(HomeQ2TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ2TotalText to Int\n", HomeQ2TotalText)
			log.Fatalln(err)
		}

		//HomeQ3Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 4)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ3TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ3TotalText, err = sportsreferenceutil.ReturnUnemptyField(HomeQ3TotalText, location, "HomeQ3Total")
		if err != nil {
			log.Fatalln(err)
		}
		matchup.HomeQ3Total, err = util.TextToInt(HomeQ3TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ3TotalText to Int\n", HomeQ3TotalText)
			log.Fatalln(err)
		}

		//HomeQ4Total
		selector = fmt.Sprintf(matchupsHomeQuarterScoresSelector, 5)
		location = fmt.Sprintf("%s %s %s", matchupsGameSummariesSelector, matchupsGameSummarySelector, selector)
		HomeQ4TotalText := util.CleanTextDatum(s.Find(selector).Text())
		HomeQ4TotalText, err = sportsreferenceutil.ReturnUnemptyField(HomeQ4TotalText, location, "HomeQ4Total")
		matchup.HomeQ4Total, err = util.TextToInt(HomeQ4TotalText)
		if err != nil {
			log.Printf("Cannot Convert '%s' for HomeQ4TotalText to Int\n", HomeQ4TotalText)
			log.Fatalln(err)
		}

		matchups = append(matchups, matchup)

	})

	if len(matchups) == 0 {
		fmt.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		fmt.Printf("Scraping of %s Completed in %s\n", url, diff)
	}
	return matchups
}
