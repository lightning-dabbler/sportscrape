package nba

import (
	"fmt"
	"log"
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
	// Selector for Basic Box Score stats tables
	basicBoxScoreSelector = `table[id$='game-basic']`
)

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
	dr := request.NewDocumentRetriever(request.WithTimeout(2 * time.Minute))
	doc, err := dr.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(basicBoxScoreSelector).Each(func(i int, s *goquery.Selection) {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).Each(func(_ int, s *goquery.Selection) {
			starterHeader = util.CleanTextDatum(s.Text())
			_, ok := basicBoxScoreStarterHeaderValues[starterHeader]
			if !ok {
				log.Fatalf("%s is not a valid Starters Header @ %s\n", starterHeader, url)
			}
		})

		s.Find(boxScoreStatsRecordsSelector).Each(func(j int, s *goquery.Selection) {
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
				boxScoreStats.PlayerLink = basketballreference.URL + util.CleanTextDatum(s.Find(boxScorePlayerLinkSelector).AttrOr("href", ""))
				boxScoreStats.Player = util.CleanTextDatum(s.Find(boxScorePlayerSelector).Text())
				playerID, err := extractPlayerID(boxScoreStats.PlayerLink)
				if err != nil {
					log.Fatalln(err)
				}
				boxScoreStats.PlayerID = playerID
				minutesPlayed := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
				if len(minutesPlayed) > 0 && unicode.IsDigit(rune(minutesPlayed[0])) {
					totalMinutes, err := transformMinutesPlayed(minutesPlayed)
					if err != nil {
						log.Fatalln(err)
					}
					boxScoreStats.MinutesPlayed = totalMinutes

					fieldGoalsMadeText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					fieldGoalsMade, err := util.TextToInt(fieldGoalsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for fieldGoalsMadeText to Int - %w; defaulting to 0.", fieldGoalsMadeText, err))
						fieldGoalsMade = 0
					}
					fieldGoalAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					fieldGoalAttempts, err := util.TextToInt(fieldGoalAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for fieldGoalAttemptsText to Int - %w; defaulting to 0.", fieldGoalAttemptsText, err))
						fieldGoalAttempts = 0
					}
					rawFieldGoalPercentage := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
					var fieldGoalPercentage float32

					if rawFieldGoalPercentage == "" {
						fieldGoalPercentage = 0
					} else {
						fieldGoalPercentage, err = util.TextToFloat32(rawFieldGoalPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Can't convert '%s' for rawFieldGoalPercentage to Float64 - %w", rawFieldGoalPercentage, err))
						}
					}

					threePointsMadeText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					threePointsMade, err := util.TextToInt(threePointsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointsMadeText to Int - %w; defaulting to 0.", threePointsMadeText, err))
						threePointsMade = 0
					}
					threePointAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					threePointAttempts, err := util.TextToInt(threePointAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointAttemptsText to Int - %w; defaulting to 0.", threePointAttemptsText, err))
						threePointsMade = 0
					}
					rawthreePointPercentage := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					var threePointPercentage float32

					if rawthreePointPercentage == "" {
						threePointPercentage = 0
					} else {
						threePointPercentage, err = util.TextToFloat32(rawthreePointPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Can't convert '%s' for rawthreePointPercentage to Float64 - %w", rawthreePointPercentage, err))
						}
					}

					freeThrowsMadeText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					freeThrowsMade, err := util.TextToInt(freeThrowsMadeText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowsMadeText to Int - %w; defaulting to 0.", freeThrowsMadeText, err))
						freeThrowsMade = 0
					}
					freeThrowAttemptsText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					freeThrowAttempts, err := util.TextToInt(freeThrowAttemptsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowAttemptsText to Int - %w; defaulting to 0.", freeThrowAttemptsText, err))
						freeThrowAttempts = 0
					}
					rawfreeThrowPercentage := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
					var freeThrowPercentage float32

					if rawfreeThrowPercentage == "" {
						freeThrowPercentage = 0
					} else {
						freeThrowPercentage, err = util.TextToFloat32(rawfreeThrowPercentage)
						if err != nil {
							log.Fatalln(fmt.Errorf("Error: Can't convert '%s' for rawfreeThrowPercentage to Float64 - %w", rawfreeThrowPercentage, err))
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
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for OffensiveReboundsText to Int - %w; defaulting to 0.", OffensiveReboundsText, err))
						boxScoreStats.OffensiveRebounds = 0
					}

					DefensiveReboundsText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					boxScoreStats.DefensiveRebounds, err = util.TextToInt(DefensiveReboundsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for DefensiveReboundsText to Int - %w; defaulting to 0.", DefensiveReboundsText, err))
						boxScoreStats.DefensiveRebounds = 0
					}

					TotalReboundsText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					boxScoreStats.TotalRebounds, err = util.TextToInt(TotalReboundsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TotalReboundsText to Int - %w; defaulting to 0.", TotalReboundsText, err))
						boxScoreStats.TotalRebounds = 0
					}

					AssistsText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					boxScoreStats.Assists, err = util.TextToInt(AssistsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for AssistsText to Int - %w; defaulting to 0.", AssistsText, err))
						boxScoreStats.Assists = 0
					}

					StealsText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					boxScoreStats.Steals, err = util.TextToInt(StealsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for StealsText to Int - %w; defaulting to 0.", StealsText, err))
						boxScoreStats.Steals = 0
					}

					BlocksText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxScoreStats.Blocks, err = util.TextToInt(BlocksText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for BlocksText to Int - %w; defaulting to 0.", BlocksText, err))
						boxScoreStats.Blocks = 0
					}

					TurnoversText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
					boxScoreStats.Turnovers, err = util.TextToInt(TurnoversText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for TurnoversText to In - %w; defaulting to 0.", TurnoversText, err))
						boxScoreStats.Turnovers = 0
					}

					PersonalFoulsText := util.CleanTextDatum(s.Find("td:nth-child(19)").Text())
					boxScoreStats.PersonalFouls, err = util.TextToInt(PersonalFoulsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PersonalFoulsText to Int - %w; defaulting to 0.", PersonalFoulsText, err))
						boxScoreStats.PersonalFouls = 0
					}

					PointsText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
					boxScoreStats.Points, err = util.TextToInt(PointsText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PointsText to Int - %w; defaulting to 0.", PointsText, err))
						boxScoreStats.Points = 0
					}

					GameScoreText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
					boxScoreStats.GameScore, err = util.TextToFloat32(GameScoreText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for GameScoreText to Float64 - %w; defaulting to 0.", GameScoreText, err))
						boxScoreStats.GameScore = 0
					}

					PlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
					boxScoreStats.PlusMinus, err = util.TextToInt(PlusMinusText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Cannot convert '%s' for PlusMinusText to Int - %w; defaulting to 0.", PlusMinusText, err))
						boxScoreStats.PlusMinus = 0
					}
				} else {
					boxScoreStats.MinutesPlayed = 0
				}
				basicNBABoxScoreStats = append(basicNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(boxScoreReserveHeaders).Each(func(_ int, s *goquery.Selection) {

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
