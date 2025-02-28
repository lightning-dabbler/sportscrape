package nba

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"

	"github.com/PuerkitoBio/goquery"
)

const (
	// https://www.basketball-reference.com/boxscores/{event_id}.html Selectors for Box Scores
	basicBoxScoreSelector             = `table[id$='game-basic']`
	basicBoxScoreContentSelector      = `#content`
	basicBoxScoreStatsRecordsSelector = `tbody > tr`
	basicBoxScoreStarterHeaders       = `thead > tr:nth-child(2) th`
	basicBoxScoreReserveHeaders       = `th`
	basicBoxScorePlayerSelector       = "th"
	basicBoxScorePlayerLinkSelector   = basicBoxScorePlayerSelector + " > a"
)

type headerValues map[string]struct{}

// basicBoxScoreStarterHeaderValues represents the headers in sequential order for the starter team members
var basicBoxScoreStarterHeaderValues headerValues = headerValues{
	"Starters": struct{}{},
	"MP":       struct{}{},
	"FG":       struct{}{},
	"FGA":      struct{}{},
	"FG%":      struct{}{},
	"3P":       struct{}{},
	"3PA":      struct{}{},
	"3P%":      struct{}{},
	"FT":       struct{}{},
	"FTA":      struct{}{},
	"FT%":      struct{}{},
	"ORB":      struct{}{},
	"DRB":      struct{}{},
	"TRB":      struct{}{},
	"AST":      struct{}{},
	"STL":      struct{}{},
	"BLK":      struct{}{},
	"TOV":      struct{}{},
	"PF":       struct{}{},
	"PTS":      struct{}{},
	"GmSc":     struct{}{},
	"+/-":      struct{}{},
}

// basicBoxScoreReservesHeaderValues represents the headers in sequential order for the reserve team members
var basicBoxScoreReservesHeaderValues headerValues = headerValues{
	"Reserves": struct{}{},
	"MP":       struct{}{},
	"FG":       struct{}{},
	"FGA":      struct{}{},
	"FG%":      struct{}{},
	"3P":       struct{}{},
	"3PA":      struct{}{},
	"3P%":      struct{}{},
	"FT":       struct{}{},
	"FTA":      struct{}{},
	"FT%":      struct{}{},
	"ORB":      struct{}{},
	"DRB":      struct{}{},
	"TRB":      struct{}{},
	"AST":      struct{}{},
	"STL":      struct{}{},
	"BLK":      struct{}{},
	"TOV":      struct{}{},
	"PF":       struct{}{},
	"PTS":      struct{}{},
	"GmSc":     struct{}{},
	"+/-":      struct{}{},
}

// getBasicBoxScoreStats accepts an interface that represents model.NBAMatchup
// It fetches the basic box score stats associated with that matchup
// Returns an array of model.NBABasicBoxScoreStats in the form of interface{}
func getBasicBoxScoreStats(nbaMatchup interface{}) []interface{} {
	matchup := nbaMatchup.(model.NBAMatchup)
	url := matchup.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var basicNBABoxScoreStats []interface{}
	fmt.Println("Scraping Basic Box Score: " + url)
	dr := request.NewDocumentRetriever(2 * time.Minute)
	doc, err := dr.RetrieveDocument(url, networkHeaders, basicBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(basicBoxScoreSelector).Each(func(i int, s *goquery.Selection) {
		var starterHeader string
		var reserveHeader string
		s.Find(basicBoxScoreStarterHeaders).Each(func(_ int, s *goquery.Selection) {

			starterHeader = util.CleanTextDatum(s.Text())
			_, ok := basicBoxScoreStarterHeaderValues[starterHeader]
			if !ok {
				log.Fatalf("%s is not a valid Starters Header @ %s\n", starterHeader, url)
			}

		})

		s.Find(basicBoxScoreStatsRecordsSelector).Each(func(j int, s *goquery.Selection) {
			var boxScoreStats model.NBABasicBoxScoreStats
			if j < 5 || j > 5 {
				boxScoreStats.PullTimestamp = PullTimestamp
				boxScoreStats.EventID = matchup.EventID
				if i == 0 {
					boxScoreStats.Team = matchup.AwayTeam
					boxScoreStats.Opponent = matchup.HomeTeam
				} else {
					boxScoreStats.Team = matchup.HomeTeam
					boxScoreStats.Opponent = matchup.AwayTeam
				}
				boxScoreStats.EventDate = matchup.EventDate
				if j < 5 {
					boxScoreStats.Starter = true
				} else {
					boxScoreStats.Starter = false
				}
				boxScoreStats.PlayerLink = basketballreference.URL + util.CleanTextDatum(s.Find(basicBoxScorePlayerLinkSelector).AttrOr("href", ""))
				boxScoreStats.Player = util.CleanTextDatum(s.Find(basicBoxScorePlayerSelector).Text())
				playerLinkSplit := strings.Split(boxScoreStats.PlayerLink, "/")
				boxScoreStats.PlayerID = strings.Split(playerLinkSplit[len(playerLinkSplit)-1], ".")[0]
				minutesPlayed := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
				if len(minutesPlayed) > 0 && unicode.IsDigit(rune(minutesPlayed[0])) {
					minutesPlayedSplit := strings.Split(minutesPlayed, ":")
					minutes, err := util.TextToInt(minutesPlayedSplit[0])
					if err != nil {
						log.Printf("Could not convert minutes %s to integer\n", minutesPlayedSplit[0])
						log.Fatalln(err)
					}

					seconds, err := util.TextToInt(minutesPlayedSplit[1])
					if err != nil {
						log.Printf("Could not convert seconds %s to integer\n", minutesPlayedSplit[0])
						log.Fatalln(err)
					}

					totalMinutes := float32(minutes) + (float32(seconds) / float32(60))
					boxScoreStats.MinutesPlayed = totalMinutes

					fieldGoalsMadeText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					fieldGoalsMade, err := util.TextToInt(fieldGoalsMadeText)
					if err != nil {
						log.Printf("Can't convert '%s' for fieldGoalsMadeText to Int\n", fieldGoalsMadeText)
						fieldGoalsMade = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					fieldGoalAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					fieldGoalAttempts, err := util.TextToInt(fieldGoalAttemptsText)
					if err != nil {
						log.Printf("Can't convert '%s' for fieldGoalAttemptsText to Int\n", fieldGoalAttemptsText)
						fieldGoalAttempts = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					rawFieldGoalPercentage := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
					var fieldGoalPercentage float32

					if rawFieldGoalPercentage == "" {
						fieldGoalPercentage = 0.00
					} else {
						fieldGoalPercentage, err = util.TextToFloat32(rawFieldGoalPercentage)
						if err != nil {
							log.Printf("Can't convert '%s' for rawFieldGoalPercentage to Float64\n", rawFieldGoalPercentage)
							log.Fatalln(err)
						}
					}

					threePointsMadeText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					threePointsMade, err := util.TextToInt(threePointsMadeText)
					if err != nil {
						log.Printf("Can't convert '%s' for threePointsMadeText to Int\n", threePointsMadeText)
						threePointsMade = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					threePointAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					threePointAttempts, err := util.TextToInt(threePointAttemptsText)
					if err != nil {
						log.Printf("Can't convert '%s' for threePointAttemptsText to Int\n", threePointAttemptsText)
						threePointsMade = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					rawthreePointPercentage := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					var threePointPercentage float32

					if rawthreePointPercentage == "" {
						threePointPercentage = 0.00
					} else {
						threePointPercentage, err = util.TextToFloat32(rawthreePointPercentage)
						if err != nil {
							log.Printf("Can't convert '%s' for rawthreePointPercentage to Float64\n", rawthreePointPercentage)
							log.Fatalln(err)
						}
					}

					freeThrowsMadeText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					freeThrowsMade, err := util.TextToInt(freeThrowsMadeText)
					if err != nil {
						log.Printf("Can't convert '%s' for freeThrowsMadeText to Int\n", freeThrowsMadeText)
						freeThrowsMade = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					freeThrowAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					freeThrowAttempts, err := util.TextToInt(freeThrowAttemptsText)
					if err != nil {
						log.Printf("Can't convert '%s' for freeThrowAttemptsText to Int\n", freeThrowAttemptsText)
						freeThrowAttempts = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
					rawfreeThrowPercentage := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
					var freeThrowPercentage float32

					if rawfreeThrowPercentage == "" {
						freeThrowPercentage = 0.00
					} else {
						freeThrowPercentage, err = util.TextToFloat32(rawfreeThrowPercentage)
						if err != nil {
							log.Printf("Can't convert '%s' for rawfreeThrowPercentage to Float64\n", rawfreeThrowPercentage)
							log.Fatalln(err)
						}
					}
					boxScoreStats.FieldGoalsMade = fieldGoalsMade
					boxScoreStats.FieldGoalAttempts = fieldGoalAttempts
					boxScoreStats.FieldGoalPercentage = fieldGoalPercentage
					boxScoreStats.ThreePointsMade = threePointsMade
					boxScoreStats.ThreePointAttempts = threePointAttempts
					boxScoreStats.ThreePointPercentage = threePointPercentage
					boxScoreStats.FreeThrowsMade = freeThrowsMade
					boxScoreStats.FreeThrowAttempts = freeThrowAttempts
					boxScoreStats.FreeThrowPercentage = freeThrowPercentage

					OffensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
					boxScoreStats.OffensiveRebounds, err = util.TextToInt(OffensiveReboundsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for OffensiveReboundsText to Int\n", OffensiveReboundsText)
						boxScoreStats.OffensiveRebounds = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					DefensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					boxScoreStats.DefensiveRebounds, err = util.TextToInt(DefensiveReboundsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for DefensiveReboundsText to Int\n", DefensiveReboundsText)
						boxScoreStats.DefensiveRebounds = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					TotalReboundsText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					boxScoreStats.TotalRebounds, err = util.TextToInt(TotalReboundsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for TotalReboundsText to Int\n", TotalReboundsText)
						boxScoreStats.TotalRebounds = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					AssistsText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					boxScoreStats.Assists, err = util.TextToInt(AssistsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for AssistsText to Int\n", AssistsText)
						boxScoreStats.Assists = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					StealsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					boxScoreStats.Steals, err = util.TextToInt(StealsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for StealsText to Int\n", StealsText)
						boxScoreStats.Steals = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					BlocksText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxScoreStats.Blocks, err = util.TextToInt(BlocksText)
					if err != nil {
						log.Printf("Cannot convert '%s' for BlocksText to Int\n", BlocksText)
						boxScoreStats.Blocks = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					TurnoversText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
					boxScoreStats.Turnovers, err = util.TextToInt(TurnoversText)
					if err != nil {
						log.Printf("Cannot convert '%s' for TurnoversText to Int\n", TurnoversText)
						boxScoreStats.Turnovers = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					PersonalFoulsText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
					boxScoreStats.PersonalFouls, err = util.TextToInt(PersonalFoulsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for PersonalFoulsText to Int\n", PersonalFoulsText)
						boxScoreStats.PersonalFouls = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					PointsText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
					boxScoreStats.Points, err = util.TextToInt(PointsText)
					if err != nil {
						log.Printf("Cannot convert '%s' for PointsText to Int\n", PointsText)
						boxScoreStats.Points = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					GameScoreText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
					boxScoreStats.GameScore, err = util.TextToFloat32(GameScoreText)
					if err != nil {
						log.Printf("Cannot convert '%s' for GameScoreText to Float64\n", GameScoreText)
						boxScoreStats.GameScore = 0
						log.Printf("WARNING: %s\n", err.Error())
					}

					PlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
					boxScoreStats.PlusMinus, err = util.TextToInt(PlusMinusText)
					if err != nil {
						log.Printf("Cannot convert '%s' for PlusMinusText to Int\n", PlusMinusText)
						boxScoreStats.PlusMinus = 0
						log.Printf("WARNING: %s\n", err.Error())
					}
				} else {
					boxScoreStats.MinutesPlayed = 0
				}
				basicNBABoxScoreStats = append(basicNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(basicBoxScoreReserveHeaders).Each(func(_ int, s *goquery.Selection) {

					reserveHeader = util.CleanTextDatum(s.Text())
					_, ok := basicBoxScoreReservesHeaderValues[reserveHeader]
					if !ok {
						log.Fatalf("%s is not a valid Reserves Header @ %s\n", reserveHeader, url)
					}

				})
			}

		})
	})
	if len(basicNBABoxScoreStats) == 0 {
		fmt.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		fmt.Printf("Scraping of %s Completed in %s\n", url, diff)
	}

	return basicNBABoxScoreStats
}

// basicBoxScoreStatsWorker acts a worker for retrieving and constructing basic box score stats
func basicBoxScoreStatsWorker(wg *sync.WaitGroup, workerNBAMatchups <-chan interface{}, boxScoreStats chan<- []interface{}) {
	for matchup := range workerNBAMatchups {
		boxScoreStats <- getBasicBoxScoreStats(matchup)
		wg.Done()
	}
}

// GetBasicBoxScoreStats retrieves all basic box score stats for all matchups
// It accepts a concurrency integer to parallelize the requests
// Returns an array of model.NBABasicBoxScoreStats in the form of interface{}
func GetBasicBoxScoreStats(concurrency int, matchups ...interface{}) []interface{} {
	var wg sync.WaitGroup
	workerNBAMatchups := make(chan interface{}, concurrency)
	BoxScoreStats := make(chan []interface{}, len(matchups))
	for i := 0; i < cap(workerNBAMatchups); i++ {
		go basicBoxScoreStatsWorker(&wg, workerNBAMatchups, BoxScoreStats)
	}
	for _, matchup := range matchups {
		wg.Add(1)
		workerNBAMatchups <- matchup
	}
	wg.Wait()
	close(workerNBAMatchups)
	close(BoxScoreStats)

	var allBoxScoreStats []interface{}
	for boxScoreStats := range BoxScoreStats {
		allBoxScoreStats = append(allBoxScoreStats, boxScoreStats...)

	}
	return allBoxScoreStats
}
