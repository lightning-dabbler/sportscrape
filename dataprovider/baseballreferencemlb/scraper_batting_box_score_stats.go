package baseballreferencemlb

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
	"github.com/xitongsys/parquet-go/types"
)

var battingBoxScoreHeaders sportsreference.Headers = sportsreference.Headers{
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

// BattingBoxScoreOption defines a configuration option for scraper
type BattingBoxScoreOption func(*BattingBoxScoreScraper)

// WithBattingBoxScoreTimeout sets the timeout duration for scraper
func WithBattingBoxScoreTimeout(timeout time.Duration) BattingBoxScoreOption {
	return func(bsr *BattingBoxScoreScraper) {
		bsr.Timeout = timeout
	}
}

// WithBattingBoxScoreDebug enables or disables debug mode for scraper
func WithBattingBoxScoreDebug(debug bool) BattingBoxScoreOption {
	return func(bsr *BattingBoxScoreScraper) {
		bsr.Debug = debug
	}
}

// NewBattingBoxScoreScraper creates a new BattingBoxScoreScraper with the provided options
func NewBattingBoxScoreScraper(options ...BattingBoxScoreOption) *BattingBoxScoreScraper {
	bsr := &BattingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(bsr)
	}
	bsr.Init()

	return bsr
}

// BattingBoxScoreScraper specialized scraper for retrieving MLB batting box score statistics
// with support for concurrent processing.
type BattingBoxScoreScraper struct {
	EventDataScraper
}

func (s BattingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballReferenceMLBBattingBoxScore
}

// Scrape retrieves MLB batting box score statistics for a single matchup.
func (s *BattingBoxScoreScraper) Scrape(matchup model.MLBMatchup) sportscrape.EventDataOutput[model.MLBBattingBoxScoreStats] {
	context := s.ConstructContext(matchup)
	output := sportscrape.EventDataOutput[model.MLBBattingBoxScoreStats]{
		Context: context,
	}

	url := matchup.BoxScoreLink
	PullTimestamp := time.Now().UTC()
	start := time.Now().UTC()
	var boxScoreStats []model.MLBBattingBoxScoreStats
	log.Println("Scraping batting Box Score: " + url)
	doc, err := s.RetrieveDocument(url, networkHeaders, contentReadySelector)
	if err != nil {
		output.Error = err
		return output
	}
	homeTeamSelector := generateStatTableSelector(matchup.HomeTeam, Batting)
	awayTeamSelector := generateStatTableSelector(matchup.AwayTeam, Batting)
	// Home team batting stat box
	homeStatBox := doc.Find(homeTeamSelector)
	// Away team batting stat box
	awayStatBox := doc.Find(awayTeamSelector)
	// Validate headers and their positions for each team's stat box
	homeStatBox.Find(headersSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := battingBoxScoreHeaders[idx]
		if header != expectedHeader {
			err = fmt.Errorf("home team header '%s' at position %d does not equal expected header '%s' @ %s", header, idx, expectedHeader, url)
			output.Error = err
			return false
		}
		return true
	})

	if output.Error != nil {
		return output
	}

	awayStatBox.Find(headersSelector).EachWithBreak(func(idx int, s *goquery.Selection) bool {
		header := util.CleanTextDatum(s.Text())
		expectedHeader := battingBoxScoreHeaders[idx]
		if header != expectedHeader {
			err = fmt.Errorf("away team header '%s' at position %d does not equal expected header '%s' @ %s", header, idx, expectedHeader, url)
			output.Error = err
			return false
		}
		return true
	})

	if output.Error != nil {
		return output
	}

	// Parse records - home team
	homeStatBox.Find(recordSelector).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		statline, err := parseBattingBoxScore(s)
		if err != nil {
			output.Error = err
			return false
		}
		if statline == nil {
			return false
		}
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchup.EventID
		statline.Team = matchup.HomeTeam
		statline.TeamID = matchup.HomeTeamID
		statline.Opponent = matchup.AwayTeam
		statline.OpponentID = matchup.AwayTeamID
		statline.EventDate = matchup.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchup.EventDate)
		boxScoreStats = append(boxScoreStats, *statline)
		return true
	})
	if output.Error != nil {
		return output
	}
	// Parse records - away team
	awayStatBox.Find(recordSelector).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		statline, err := parseBattingBoxScore(s)
		if err != nil {
			output.Error = err
			return false
		}
		if statline == nil {
			return false
		}
		statline.PullTimestamp = PullTimestamp
		statline.EventID = matchup.EventID
		statline.Team = matchup.AwayTeam
		statline.TeamID = matchup.AwayTeamID
		statline.Opponent = matchup.HomeTeam
		statline.OpponentID = matchup.HomeTeamID
		statline.EventDate = matchup.EventDate
		statline.PullTimestampParquet = types.TimeToTIMESTAMP_MILLIS(PullTimestamp, true)
		statline.EventDateParquet = util.TimeToDays(matchup.EventDate)
		boxScoreStats = append(boxScoreStats, *statline)
		return true
	})
	if output.Error != nil {
		return output
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of %s Completed in %s\n", url, diff)
	output.Output = boxScoreStats
	return output
}

// parseBattingBoxScore parses a player's batting statline
//
// Parameters:
//   - s: goquery.Selection representing a player's batting statline
func parseBattingBoxScore(s *goquery.Selection) (*model.MLBBattingBoxScoreStats, error) {
	if s.AttrOr("class", "") == "spacer" {
		return nil, nil
	}
	var statline model.MLBBattingBoxScoreStats
	// Position
	position := util.CleanTextDatum(s.Find(positionSelector).Text())
	positionSplit := strings.Split(position, " ")
	statline.Position = positionSplit[len(positionSplit)-1]
	// Player, PlayerLink, & PlayerID
	player := s.Find(playerSelector)
	statline.PlayerLink = sportsreference.BaseballRefURL + util.CleanTextDatum(player.AttrOr("href", ""))
	statline.Player = util.CleanTextDatum(player.Text())
	playerID, err := sportsreference.PlayerID(statline.PlayerLink)
	if err != nil {
		return nil, err
	}
	statline.PlayerID = playerID
	// AtBat
	atBatText := util.CleanTextDatum(s.Find("td:nth-child(2)").Text())
	atBat, err := util.TextToInt32(atBatText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for atBatText to int - %w", atBatText, err)
		return nil, err
	}
	statline.AtBat = atBat
	// Runs
	runsText := util.CleanTextDatum(s.Find("td:nth-child(3)").Text())
	runs, err := util.TextToInt32(runsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for runsText to int - %w", runsText, err)
		return nil, err
	}
	statline.Runs = runs
	// Hits
	hitsText := util.CleanTextDatum(s.Find("td:nth-child(4)").Text())
	hits, err := util.TextToInt32(hitsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for hitsText to int - %w", hitsText, err)
		return nil, err
	}
	statline.Hits = hits
	// RunsBattedIn
	runsBattedInText := util.CleanTextDatum(s.Find("td:nth-child(5)").Text())
	runsBattedIn, err := util.TextToInt32(runsBattedInText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for runsBattedInText to int - %w", runsBattedInText, err)
		return nil, err
	}
	statline.RunsBattedIn = runsBattedIn
	// Walks
	walksText := util.CleanTextDatum(s.Find("td:nth-child(6)").Text())
	walks, err := util.TextToInt32(walksText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for walksText to int - %w", walksText, err)
		return nil, err
	}
	statline.Walks = walks
	// Strikeouts
	strikeoutsText := util.CleanTextDatum(s.Find("td:nth-child(7)").Text())
	strikeouts, err := util.TextToInt32(strikeoutsText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for strikeoutsText to int - %w", strikeoutsText, err)
		return nil, err
	}
	statline.Strikeouts = strikeouts
	// PlateAppearances
	plateAppearancesText := util.CleanTextDatum(s.Find("td:nth-child(8)").Text())
	plateAppearances, err := util.TextToInt32(plateAppearancesText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for plateAppearancesText to int - %w", plateAppearancesText, err)
		return nil, err
	}
	statline.PlateAppearances = plateAppearances
	// BattingAverage
	battingAverageText := util.CleanTextDatum(s.Find("td:nth-child(9)").Text())
	if battingAverageText != "" {
		battingAverage, err := util.TextToFloat32(battingAverageText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for battingAverageText to float32 - %w", battingAverageText, err)
			return nil, err
		}
		statline.BattingAverage = &battingAverage
	}
	// OnBasePercentage
	onBasePercentageText := util.CleanTextDatum(s.Find("td:nth-child(10)").Text())
	if onBasePercentageText != "" {
		onBasePercentage, err := util.TextToFloat32(onBasePercentageText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for onBasePercentageText to float32 - %w", onBasePercentageText, err)
			return nil, err
		}
		statline.OnBasePercentage = &onBasePercentage
	}
	// SluggingPercentage
	sluggingPercentageText := util.CleanTextDatum(s.Find("td:nth-child(11)").Text())
	if sluggingPercentageText != "" {
		sluggingPercentage, err := util.TextToFloat32(sluggingPercentageText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for sluggingPercentageText to float32 - %w", sluggingPercentageText, err)
			return nil, err
		}
		statline.SluggingPercentage = &sluggingPercentage
	}
	// OnBasePlusSlugging
	onBasePlusSluggingText := util.CleanTextDatum(s.Find("td:nth-child(12)").Text())
	if onBasePlusSluggingText != "" {
		onBasePlusSlugging, err := util.TextToFloat32(onBasePlusSluggingText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for onBasePlusSluggingText to float32 - %w", onBasePlusSluggingText, err)
			return nil, err
		}
		statline.OnBasePlusSlugging = &onBasePlusSlugging
	}
	// PitchesPerPlateAppearance
	pitchesPerPlateAppearanceText := util.CleanTextDatum(s.Find("td:nth-child(13)").Text())
	if pitchesPerPlateAppearanceText != "" {
		pitchesPerPlateAppearance, err := util.TextToInt32(pitchesPerPlateAppearanceText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for pitchesPerPlateAppearanceText to int - %w", pitchesPerPlateAppearanceText, err)
			return nil, err
		}
		statline.PitchesPerPlateAppearance = &pitchesPerPlateAppearance
	}
	// Strikes
	strikesText := util.CleanTextDatum(s.Find("td:nth-child(14)").Text())
	if strikesText != "" {
		strikes, err := util.TextToInt32(strikesText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for strikesText to int - %w", strikesText, err)
			return nil, err
		}
		statline.Strikes = &strikes
	}
	// WinProbabilityAdded
	winProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(15)").Text())
	if winProbabilityAddedText != "" {
		winProbabilityAdded, err := util.TextToFloat32(winProbabilityAddedText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for winProbabilityAddedText to float32 - %w", winProbabilityAddedText, err)
			return nil, err
		}
		statline.WinProbabilityAdded = &winProbabilityAdded
	}
	// AverageLeverageIndex
	averageLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(16)").Text())
	if averageLeverageIndexText != "" {
		averageLeverageIndex, err := util.TextToFloat32(averageLeverageIndexText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for averageLeverageIndexText to float32 - %w", averageLeverageIndexText, err)
			return nil, err
		}
		statline.AverageLeverageIndex = &averageLeverageIndex
	}
	// SumPositiveWinProbabilityAdded
	sumPositiveWinProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(17)").Text())
	if sumPositiveWinProbabilityAddedText != "" {
		sumPositiveWinProbabilityAdded, err := util.TextToFloat32(sumPositiveWinProbabilityAddedText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for sumPositiveWinProbabilityAddedText to float32 - %w", sumPositiveWinProbabilityAddedText, err)
			return nil, err
		}
		statline.SumPositiveWinProbabilityAdded = &sumPositiveWinProbabilityAdded
	}
	// SumNegativeWinProbabilityAdded
	sumNegativeWinProbabilityAddedText := util.CleanTextDatum(s.Find("td:nth-child(18)").Text())
	if sumNegativeWinProbabilityAddedText != "" {
		sumNegativeWinProbabilityAdded, err := util.TextToFloat32(sumNegativeWinProbabilityAddedText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for sumNegativeWinProbabilityAddedText to float32 - %w", sumNegativeWinProbabilityAddedText, err)
			return nil, err
		}
		statline.SumNegativeWinProbabilityAdded = &sumNegativeWinProbabilityAdded
	}
	// ChampionshipWinProbabilityAdded
	championshipWinProbabilityAddedText := strings.TrimRight(util.CleanTextDatum(s.Find("td:nth-child(19)").Text()), "%") // 0.13% --> 0.13
	if championshipWinProbabilityAddedText != "" {
		championshipWinProbabilityAdded, err := util.TextToFloat32(championshipWinProbabilityAddedText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for championshipWinProbabilityAddedText to float32 - %w", championshipWinProbabilityAddedText, err)
			return nil, err
		}
		statline.ChampionshipWinProbabilityAdded = &championshipWinProbabilityAdded
	}
	// AverageChampionshipLeverageIndex
	averageChampionshipLeverageIndexText := util.CleanTextDatum(s.Find("td:nth-child(20)").Text())
	if averageChampionshipLeverageIndexText != "" {
		averageChampionshipLeverageIndex, err := util.TextToFloat32(averageChampionshipLeverageIndexText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for averageChampionshipLeverageIndexText to float32 - %w", averageChampionshipLeverageIndexText, err)
			return nil, err
		}
		statline.AverageChampionshipLeverageIndex = &averageChampionshipLeverageIndex
	}
	// BaseOutRunsAdded
	baseOutRunsAddedText := util.CleanTextDatum(s.Find("td:nth-child(21)").Text())
	if baseOutRunsAddedText != "" {
		baseOutRunsAdded, err := util.TextToFloat32(baseOutRunsAddedText)
		if err != nil {
			err = fmt.Errorf("ERROR: Can't convert '%s' for baseOutRunsAddedText to float32 - %w", baseOutRunsAddedText, err)
			return nil, err
		}
		statline.BaseOutRunsAdded = &baseOutRunsAdded
	}
	// Putout
	putoutText := util.CleanTextDatum(s.Find("td:nth-child(22)").Text())
	putout, err := util.TextToInt32(putoutText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for putoutText to int - %w", putoutText, err)
		return nil, err
	}
	statline.Putout = putout
	// Assist
	assistText := util.CleanTextDatum(s.Find("td:nth-child(23)").Text())
	assist, err := util.TextToInt32(assistText)
	if err != nil {
		err = fmt.Errorf("ERROR: Can't convert '%s' for assistText to int - %w", assistText, err)
		return nil, err
	}
	statline.Assist = assist
	return &statline, nil
}
