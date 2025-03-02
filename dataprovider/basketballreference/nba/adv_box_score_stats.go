package nba

import (
	"fmt"
	"log"
	"sync"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"
)

const (
	// Selector for Advanced Box Score stats tables
	advBoxScoreSelector = `table[id$='game-advanced']`
)

// advBoxScoreStarterHeaderValues represents the headers in sequential order for the starter team members
var advBoxScoreStarterHeaderValues headerValues = headerValues{
	"Starters": struct{}{},
	"MP":       struct{}{},
	"TS%":      struct{}{},
	"eFG%":     struct{}{},
	"3PAr":     struct{}{},
	"FTr":      struct{}{},
	"ORB%":     struct{}{},
	"DRB%":     struct{}{},
	"TRB%":     struct{}{},
	"AST%":     struct{}{},
	"STL%":     struct{}{},
	"BLK%":     struct{}{},
	"TOV%":     struct{}{},
	"USG%":     struct{}{},
	"ORtg":     struct{}{},
	"DRtg":     struct{}{},
	"BPM":      struct{}{},
}

// advBoxScoreReservesHeaderValues represents the headers in sequential order for the reserve team members
var advBoxScoreReservesHeaderValues headerValues = headerValues{
	"Reserves": struct{}{},
	"MP":       struct{}{},
	"TS%":      struct{}{},
	"eFG%":     struct{}{},
	"3PAr":     struct{}{},
	"FTr":      struct{}{},
	"ORB%":     struct{}{},
	"DRB%":     struct{}{},
	"TRB%":     struct{}{},
	"AST%":     struct{}{},
	"STL%":     struct{}{},
	"BLK%":     struct{}{},
	"TOV%":     struct{}{},
	"USG%":     struct{}{},
	"ORtg":     struct{}{},
	"DRtg":     struct{}{},
	"BPM":      struct{}{},
}

// getAdvBoxScoreStats accepts an interface that represents model.NBAMatchup
// It fetches the advanced box score stats associated with that matchup
// Returns an array of model.NBAAdvBoxScoreStats in the form of interface{}
func getAdvBoxScoreStats(nbaMatchup interface{}) []interface{} {
	matchup := nbaMatchup.(model.NBAMatchup)
	url := matchup.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var advNBABoxScoreStats []interface{}
	fmt.Println("Scraping Advanced Box Score: " + url)
	dr := request.NewDocumentRetriever(request.WithTimeout(2 * time.Minute))
	doc, err := dr.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(advBoxScoreSelector).Each(func(i int, s *goquery.Selection) {
		var starterHeader string
		var reserveHeader string
		s.Find(boxScoreStarterHeaders).Each(func(_ int, s *goquery.Selection) {

			starterHeader = util.CleanTextDatum(s.Text())
			_, ok := advBoxScoreStarterHeaderValues[starterHeader]
			if !ok {
				log.Fatalf("%s is not a valid Starters Header @ %s\n", starterHeader, url)
			}

		})

		s.Find(boxScoreStatsRecordsSelector).Each(func(j int, s *goquery.Selection) {
			var boxScoreStats model.NBAAdvBoxScoreStats
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
					// TrueShootingPercentage
					trueShootingPercentageText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
					trueShootingPercentage, err := util.TextToFloat32(trueShootingPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for trueShootingPercentageText to float32 - %w; defaulting to 0.", trueShootingPercentageText, err))
						trueShootingPercentage = 0
					}
					boxScoreStats.TrueShootingPercentage = trueShootingPercentage

					// EffectiveFieldGoalPercentage
					effectiveFieldGoalPercentageText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
					effectiveFieldGoalPercentage, err := util.TextToFloat32(effectiveFieldGoalPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for effectiveFieldGoalPercentageText to float32 - %w; defaulting to 0.", effectiveFieldGoalPercentageText, err))
						effectiveFieldGoalPercentage = 0
					}
					boxScoreStats.EffectiveFieldGoalPercentage = effectiveFieldGoalPercentage

					// ThreePointAttemptRate
					threePointAttemptRateText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
					threePointAttemptRate, err := util.TextToFloat32(threePointAttemptRateText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for threePointAttemptRateText to float32 - %w; defaulting to 0.", threePointAttemptRateText, err))
						threePointAttemptRate = 0
					}
					boxScoreStats.ThreePointAttemptRate = threePointAttemptRate

					// FreeThrowAttemptRate
					freeThrowAttemptRateText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
					freeThrowAttemptRate, err := util.TextToFloat32(freeThrowAttemptRateText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for freeThrowAttemptRateText to float32 - %w; defaulting to 0.", freeThrowAttemptRateText, err))
						freeThrowAttemptRate = 0
					}
					boxScoreStats.FreeThrowAttemptRate = freeThrowAttemptRate

					// OffensiveReboundPercentage
					offensiveReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
					offensiveReboundPercentage, err := util.TextToFloat32(offensiveReboundPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for offensiveReboundPercentageText to float32 - %w; defaulting to 0.", offensiveReboundPercentageText, err))
						offensiveReboundPercentage = 0
					}
					boxScoreStats.OffensiveReboundPercentage = offensiveReboundPercentage

					// DefensiveReboundPercentage
					defensiveReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
					defensiveReboundPercentage, err := util.TextToFloat32(defensiveReboundPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for defensiveReboundPercentageText to float32 - %w; defaulting to 0.", defensiveReboundPercentageText, err))
						defensiveReboundPercentage = 0
					}
					boxScoreStats.DefensiveReboundPercentage = defensiveReboundPercentage

					// TotalReboundPercentage
					totalReboundPercentageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
					totalReboundPercentage, err := util.TextToFloat32(totalReboundPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for totalReboundPercentageText to float32 - %w; defaulting to 0.", totalReboundPercentageText, err))
						totalReboundPercentage = 0
					}
					boxScoreStats.TotalReboundPercentage = totalReboundPercentage

					// AssistPercentage
					assistPercentageText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
					assistPercentage, err := util.TextToFloat32(assistPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for assistPercentageText to float32 - %w; defaulting to 0.", assistPercentageText, err))
						assistPercentage = 0
					}
					boxScoreStats.AssistPercentage = assistPercentage

					// StealPercentage
					stealPercentageText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
					stealPercentage, err := util.TextToFloat32(stealPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for stealPercentageText to float32 - %w; defaulting to 0.", stealPercentageText, err))
						stealPercentage = 0
					}
					boxScoreStats.StealPercentage = stealPercentage

					// BlockPercentage
					blockPercentageText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
					blockPercentage, err := util.TextToFloat32(blockPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for blockPercentageText to float32 - %w; defaulting to 0.", blockPercentageText, err))
						blockPercentage = 0
					}
					boxScoreStats.BlockPercentage = blockPercentage

					// TurnoverPercentage
					turnoverPercentageText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
					turnoverPercentage, err := util.TextToFloat32(turnoverPercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for turnoverPercentageText to float32 - %w; defaulting to 0.", turnoverPercentageText, err))
						turnoverPercentage = 0
					}
					boxScoreStats.TurnoverPercentage = turnoverPercentage

					// UsagePercentage
					usagePercentageText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
					usagePercentage, err := util.TextToFloat32(usagePercentageText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for usagePercentageText to float32 - %w; defaulting to 0.", usagePercentageText, err))
						usagePercentage = 0
					}
					boxScoreStats.UsagePercentage = usagePercentage

					// OffensiveRating
					offensiveRatingText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
					offensiveRating, err := util.TextToInt(offensiveRatingText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for offensiveRatingText to float32 - %w; defaulting to 0.", offensiveRatingText, err))
						offensiveRating = 0
					}
					boxScoreStats.OffensiveRating = offensiveRating

					// DefensiveRating
					defensiveRatingText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
					defensiveRating, err := util.TextToInt(defensiveRatingText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for defensiveRatingText to float32 - %w; defaulting to 0.", defensiveRatingText, err))
						defensiveRating = 0
					}
					boxScoreStats.DefensiveRating = defensiveRating

					// BoxPlusMinus
					boxPlusMinusText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
					boxPlusMinus, err := util.TextToFloat32(boxPlusMinusText)
					if err != nil {
						log.Println(fmt.Errorf("WARNING: Can't convert '%s' for boxPlusMinusText to float32 - %w; defaulting to 0.", boxPlusMinusText, err))
						boxPlusMinus = 0
					}
					boxScoreStats.BoxPlusMinus = boxPlusMinus

				} else {
					boxScoreStats.MinutesPlayed = 0
				}

				advNBABoxScoreStats = append(advNBABoxScoreStats, boxScoreStats)
			} else {
				s.Find(boxScoreReserveHeaders).Each(func(_ int, s *goquery.Selection) {

					reserveHeader = util.CleanTextDatum(s.Text())
					_, ok := advBoxScoreReservesHeaderValues[reserveHeader]
					if !ok {
						log.Fatalf("%s is not a valid Reserves Header @ %s\n", reserveHeader, url)
					}
				})
			}
		})
	})
	if len(advNBABoxScoreStats) == 0 {
		fmt.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		fmt.Printf("Scraping of %s Completed in %s\n", url, diff)
	}

	return advNBABoxScoreStats
}

// advBoxScoreStatsWorker acts a worker for retrieving and constructing advanced box score stats
func advBoxScoreStatsWorker(wg *sync.WaitGroup, workerNBAMatchups <-chan interface{}, boxScoreStats chan<- []interface{}) {
	for matchup := range workerNBAMatchups {
		boxScoreStats <- getAdvBoxScoreStats(matchup)
		wg.Done()
	}
}

// GetAdvBoxScoreStats retrieves all advanced box score stats for all matchups
// It accepts a concurrency integer to parallelize the requests
// Returns an array of model.NBAAdvBoxScoreStats in the form of interface{}
func GetAdvBoxScoreStats(concurrency int, matchups ...interface{}) []interface{} {
	var wg sync.WaitGroup
	workerNBAMatchups := make(chan interface{}, concurrency)
	BoxScoreStats := make(chan []interface{}, len(matchups))
	for i := 0; i < cap(workerNBAMatchups); i++ {
		go advBoxScoreStatsWorker(&wg, workerNBAMatchups, BoxScoreStats)
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
