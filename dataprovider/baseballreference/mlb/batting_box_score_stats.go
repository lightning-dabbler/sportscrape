package mlb

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	sportsreferenceutil "github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

var battingBoxScoreHeaders sportsreferenceutil.Headers = sportsreferenceutil.Headers{
	"Batting",
	"AB",
	"R",
	"H",
	"RBI",
	"BB",
	"SO",
	"PA",
	"BA",
	"OBP",
	"SLG",
	"OPS",
	"Pit",
	"Str",
	"WPA",
	"aLI",
	"WPA+",
	"WPA-",
	"cWPA",
	"acLI",
	"RE24",
	"PO",
	"A",
	"Details",
}

// BattingBoxScoreOption defines a configuration option for batting box score runners
type BattingBoxScoreOption func(*BattingBoxScoreRunner)

// WithBattingBoxScoreTimeout sets the timeout duration for batting box score runner
func WithBattingBoxScoreTimeout(timeout time.Duration) BattingBoxScoreOption {
	return func(bsr *BattingBoxScoreRunner) {
		bsr.Timeout = timeout
	}
}

// WithBattingBoxScoreDebug enables or disables debug mode for batting box score runner
func WithBattingBoxScoreDebug(debug bool) BattingBoxScoreOption {
	return func(bsr *BattingBoxScoreRunner) {
		bsr.Debug = debug
	}
}

// WithBattingBoxScoreConcurrency sets the number of concurrent workers
func WithBattingBoxScoreConcurrency(n int) BattingBoxScoreOption {
	return func(bsr *BattingBoxScoreRunner) {
		bsr.Concurrency = n
	}
}

// NewBattingBoxScoreRunner creates a new BattingBoxScoreRunner with the provided options
func NewBattingBoxScoreRunner(options ...BattingBoxScoreOption) *BattingBoxScoreRunner {
	bsr := &BattingBoxScoreRunner{}
	bsr.Processor = bsr
	//

	// Apply all options
	for _, option := range options {
		option(bsr)
	}

	return bsr
}

// BattingBoxScoreRunner specialized Runner for retrieving MLB batting box score statistics
// with support for concurrent processing.
type BattingBoxScoreRunner struct {
	sportsreferenceutil.BoxScoreRunner
}

// GetSegmentBoxScoreStats retrieves MLB batting box score statistics for a single matchup.
//
// Parameter:
//   - matchup: The MLB matchup for which to retrieve batting box score statistics
//
// Returns a slice of MLB batting box score statistics as interface{} values
func (boxScoreRunner *BattingBoxScoreRunner) GetSegmentBoxScoreStats(matchup interface{}) []interface{} {
	matchupModel := matchup.(model.MLBMatchup)
	url := matchupModel.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var boxScoreStats []interface{}
	log.Println("Scraping batting Box Score: " + url)
	doc, err := boxScoreRunner.RetrieveDocument(url, networkHeaders, waitReadyBoxScoreContentSelector)
	if err != nil {
		log.Fatalln(err)
	}
	homeTeamSelector := generateStatTableSelector(matchupModel.HomeTeam, Batting)
	awayTeamSelector := generateStatTableSelector(matchupModel.AwayTeam, Batting)
	// Home team batting stat box
	homeStatBox := doc.Find(homeTeamSelector)
	// Away team batting stat box
	awayStatBox := doc.Find(awayTeamSelector)
	// Validate headers and their positions for each team's stat box
	homeStatBox.Find(headersSelector).Each(func(idx int, s *goquery.Selection) {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := battingBoxScoreHeaders[idx]
		if header != expectedHeader {
			log.Fatalf("Home team header '%s' at position %d does not equal expected header '%s' @ %s\n", header, idx, expectedHeader, url)
		}
	})

	awayStatBox.Find(headersSelector).Each(func(idx int, s *goquery.Selection) {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := battingBoxScoreHeaders[idx]
		if header != expectedHeader {
			log.Fatalf("Away team header '%s' at position %d does not equal expected header '%s' @ %s\n", header, idx, expectedHeader, url)
		}
	})

	// Parse records - home team
	homeStatBox.Find(recordSelector).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		statline := parseBattingBoxScore(s)
		if statline == nil {
			return false
		}
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.HomeTeam
		statline.Opponent = matchupModel.AwayTeam
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, *statline)
		return true
	})
	// Parse records - away team
	awayStatBox.Find(recordSelector).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		statline := parseBattingBoxScore(s)
		if statline == nil {
			return false
		}
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchupModel.EventID
		statline.Team = matchupModel.AwayTeam
		statline.Opponent = matchupModel.HomeTeam
		statline.EventDate = matchupModel.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchupModel.EventDate)
		boxScoreStats = append(boxScoreStats, *statline)
		return true
	})

	if len(boxScoreStats) == 0 {
		log.Printf("No Data Scraped @ %s\n", url)
	} else {
		diff := time.Now().UTC().Sub(start)
		log.Printf("Scraping of %s Completed in %s\n", url, diff)
	}
	return boxScoreStats
}

// parseBattingBoxScore parses a player's batting statline
//
// Parameters:
//   - s: goquery.Selection representing a player's batting statline
//
// Returns model.MLBBattingBoxScoreStats containing parsed statistics
func parseBattingBoxScore(s *goquery.Selection) *model.MLBBattingBoxScoreStats {
	if s.AttrOr("class", "") == "spacer" {
		return nil
	}
	var statline model.MLBBattingBoxScoreStats
	// Position
	position := util.CleanTextDatum(s.Find(positionSelector).Text())
	positionSplit := strings.Split(position, " ")
	statline.Position = positionSplit[len(positionSplit)-1]
	// Player, PlayerLink, & PlayerID
	player := s.Find(playerSelector)
	statline.PlayerLink = baseballreference.URL + util.CleanTextDatum(player.AttrOr("href", ""))
	statline.Player = util.CleanTextDatum(player.Text())
	playerID, err := sportsreferenceutil.PlayerID(statline.PlayerLink)
	if err != nil {
		log.Fatalln(err)
	}
	statline.PlayerID = playerID
	// AtBat
	atBatText := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
	atBat, err := util.TextToInt32(atBatText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for atBatText to int - %w", atBatText, err))
	}
	statline.AtBat = atBat
	// Runs
	runsText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
	runs, err := util.TextToInt32(runsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for runsText to int - %w", runsText, err))
	}
	statline.Runs = runs
	// Hits
	hitsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
	hits, err := util.TextToInt32(hitsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for hitsText to int - %w", hitsText, err))
	}
	statline.Hits = hits
	// RunsBattedIn
	runsBattedInText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
	runsBattedIn, err := util.TextToInt32(runsBattedInText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for runsBattedInText to int - %w", runsBattedInText, err))
	}
	statline.RunsBattedIn = runsBattedIn
	// Walks
	walksText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
	walks, err := util.TextToInt32(walksText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for walksText to int - %w", walksText, err))
	}
	statline.Walks = walks
	// Strikeouts
	strikeoutsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
	strikeouts, err := util.TextToInt32(strikeoutsText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikeoutsText to int - %w", strikeoutsText, err))
	}
	statline.Strikeouts = strikeouts
	// PlateAppearances
	plateAppearancesText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
	plateAppearances, err := util.TextToInt32(plateAppearancesText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for plateAppearancesText to int - %w", plateAppearancesText, err))
	}
	statline.PlateAppearances = plateAppearances
	// BattingAverage
	battingAverageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
	if battingAverageText != "" {
		battingAverage, err := util.TextToFloat32(battingAverageText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for battingAverageText to float32 - %w!", battingAverageText, err))
		}
		statline.BattingAverage = &battingAverage
	}
	// OnBasePercentage
	onBasePercentageText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
	if onBasePercentageText != "" {
		onBasePercentage, err := util.TextToFloat32(onBasePercentageText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for onBasePercentageText to float32 - %w!", onBasePercentageText, err))
		}
		statline.OnBasePercentage = &onBasePercentage
	}
	// SluggingPercentage
	sluggingPercentageText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
	if sluggingPercentageText != "" {
		sluggingPercentage, err := util.TextToFloat32(sluggingPercentageText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for sluggingPercentageText to float32 - %w!", sluggingPercentageText, err))
		}
		statline.SluggingPercentage = &sluggingPercentage
	}
	// OnBasePlusSlugging
	onBasePlusSluggingText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
	if onBasePlusSluggingText != "" {
		onBasePlusSlugging, err := util.TextToFloat32(onBasePlusSluggingText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for onBasePlusSluggingText to float32 - %w!", onBasePlusSluggingText, err))
		}
		statline.OnBasePlusSlugging = &onBasePlusSlugging
	}
	// PitchesPerPlateAppearance
	pitchesPerPlateAppearanceText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
	if pitchesPerPlateAppearanceText != "" {
		pitchesPerPlateAppearance, err := util.TextToInt32(pitchesPerPlateAppearanceText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for pitchesPerPlateAppearanceText to int - %w!", pitchesPerPlateAppearanceText, err))
		}
		statline.PitchesPerPlateAppearance = &pitchesPerPlateAppearance
	}
	// Strikes
	strikesText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
	if strikesText != "" {
		strikes, err := util.TextToInt32(strikesText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for strikesText to int - %w!", strikesText, err))
		}
		statline.Strikes = &strikes
	}
	// WinProbabilityAdded
	winProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
	if winProbabilityAddedText != "" {
		winProbabilityAdded, err := util.TextToFloat32(winProbabilityAddedText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for winProbabilityAddedText to float32 - %w!", winProbabilityAddedText, err))
		}
		statline.WinProbabilityAdded = &winProbabilityAdded
	}
	// AverageLeverageIndex
	averageLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
	if averageLeverageIndexText != "" {
		averageLeverageIndex, err := util.TextToFloat32(averageLeverageIndexText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for averageLeverageIndexText to float32 - %w!", averageLeverageIndexText, err))
		}
		statline.AverageLeverageIndex = &averageLeverageIndex
	}
	// SumPositiveWinProbabilityAdded
	sumPositiveWinProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
	if sumPositiveWinProbabilityAddedText != "" {
		sumPositiveWinProbabilityAdded, err := util.TextToFloat32(sumPositiveWinProbabilityAddedText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for sumPositiveWinProbabilityAddedText to float32 - %w!", sumPositiveWinProbabilityAddedText, err))
		}
		statline.SumPositiveWinProbabilityAdded = &sumPositiveWinProbabilityAdded
	}
	// SumNegativeWinProbabilityAdded
	sumNegativeWinProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
	if sumNegativeWinProbabilityAddedText != "" {
		sumNegativeWinProbabilityAdded, err := util.TextToFloat32(sumNegativeWinProbabilityAddedText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for sumNegativeWinProbabilityAddedText to float32 - %w!", sumNegativeWinProbabilityAddedText, err))
		}
		statline.SumNegativeWinProbabilityAdded = &sumNegativeWinProbabilityAdded
	}
	// ChampionshipWinProbabilityAdded
	championshipWinProbabilityAddedText := strings.TrimRight(util.CleanTextDatum(s.Find("td:nth-child(19)").Text()), "%") // 0.13% --> 0.13
	if championshipWinProbabilityAddedText != "" {
		championshipWinProbabilityAdded, err := util.TextToFloat32(championshipWinProbabilityAddedText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for championshipWinProbabilityAddedText to float32 - %w!", championshipWinProbabilityAddedText, err))
		}
		statline.ChampionshipWinProbabilityAdded = &championshipWinProbabilityAdded
	}
	// AverageChampionshipLeverageIndex
	averageChampionshipLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
	if averageChampionshipLeverageIndexText != "" {
		averageChampionshipLeverageIndex, err := util.TextToFloat32(averageChampionshipLeverageIndexText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for averageChampionshipLeverageIndexText to float32 - %w!", averageChampionshipLeverageIndexText, err))
		}
		statline.AverageChampionshipLeverageIndex = &averageChampionshipLeverageIndex
	}
	// BaseOutRunsAdded
	baseOutRunsAddedText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
	if baseOutRunsAddedText != "" {
		baseOutRunsAdded, err := util.TextToFloat32(baseOutRunsAddedText)
		if err != nil {
			log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for baseOutRunsAddedText to float32 - %w!", baseOutRunsAddedText, err))
		}
		statline.BaseOutRunsAdded = &baseOutRunsAdded
	}
	// Putout
	putoutText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
	putout, err := util.TextToInt32(putoutText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for putoutText to int - %w", putoutText, err))
	}
	statline.Putout = putout
	// Assist
	assistText := util.CleanTextDatum(s.Find("td:nth-child(23)").Text())
	assist, err := util.TextToInt32(assistText)
	if err != nil {
		log.Fatalln(fmt.Errorf("ERROR: Can't convert '%s' for assistText to int - %w", assistText, err))
	}
	statline.Assist = assist
	return &statline
}
