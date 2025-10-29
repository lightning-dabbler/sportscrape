# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.16.0] - 2025-10-28
### Added
- Adds a scraper for ESPN MMA data (UFC/PFL). (#105)
- Matchup scraper for scraping event data (UFC/PFL). (#105)
- Fight Details Scraper for scraping fight outcomes and stats (UFC/PFL). (#105)

### Fixed
- Fixed pitching and batting scraper tests in baseball reference (#106)

### New Contributors
- @vlaurenzano in #105

## [0.15.1] - 2025-08-29
### Fixed
-  foxsports `MLBOddsTotal.OverOdds` and `MLBOddsTotal.UnderOdds` are now optional because it's not always available for all events (#103)
## [0.15.0] - 2025-08-28
### Changed
- Chained errors observed in `EventDataRunner` and `MatchupRunner` (#99)
### Fixed
- foxsports `MLBProbableStartingPitcher.StartingPitcherRecord` and `MLBProbableStartingPitcher.StartingPitcherERA` are now nilable (#101)
- Fixed feed name for `FSMLBProbableStartingPitcher` (#101)
### Added
- Added the following fields to the baseball savant mlb `Matchup` model (#100)
    - VenueID
    - VenueName
    - HomeTeamLeagueID
    - HomeTeamLeagueName
    - HomeTeamDivisionID
    - HomeTeamDivisionName
    - HomeStartingPitcherPitchHand
    - AwayTeamLeagueID
    - AwayTeamLeagueName
    - AwayTeamDivisionID
    - AwayTeamDivisionName
    - AwayStartingPitcherPitchHand
- Added foxsports `MLBOddsTotalScraper` and `MLBOddsMoneyLineScraper` betting odds extraction (#101)
### Documentation
- foxsports `MLBProbableStartingPitcher.StartingPitcherRecord` and `MLBProbableStartingPitcher.StartingPitcherERA` are not point-in-time (PIT) with respect to event (#101)
    - Added comment to data model
- Added foxsports mlb odds total and model line to catalog in README (#101)
## [0.14.0] - 2025-08-08
### Added
- Added fox sports wnba matchup and box score feeds (#97)
### Changed
- Modified `NBABoxScoreScraper` to support WNBA extraction (#97)
- `EventDataRunner.Run()` and `MatchupRunner.Run()` now call `Scraper.Init()` for extra validation before execution (#97)
### Documentation
- Added fox sports WNBA matchup and box score to provider list (#97)

## [0.13.4] - 2025-07-29
### Fixed
- Fixed parquet field names for baseball savant mlb matchup data model

## [0.13.3] - 2025-07-27
### Fixed
- Fixed unmarshalling of baseball savant mlb matchups to handle empty response payload (#94)

## [0.13.2] - 2025-07-17
### Changed
- `FSNCAAMatchup` -> `FSNCAABMatchup` (#93)

## [0.13.1] - 2025-07-11
### Changed
- Documentation updates

## [0.13.0] - 2025-07-11
### Changed
- Moved mocks to internal (#90)
- Code organizational improvements (#90)

## [0.12.1] - 2025-07-10
### Fixed
- Fixed links that reference data models in readme (#89)

## [0.12.0] - 2025-07-10
### Added
- Added `MatchupScraper` for baseball savant mlb (#84)
- Added `FieldingBoxScoreScraper`,`BattingBoxScoreScraper`, `PitchingBoxScoreScraper`, and `PlayByPlayScraper` for baseball savant mlb (#86)
### Changed
- Moved `RFC3339ToTime` and `DateStrToTime` to time.go (#84)
- Documentation updates (#84)
- Deprecation warning on baseball reference feed (#87)

## [0.11.0] - 2025-07-02
### Added
- Added `MatchupRunner`, `EventDataRunner`, `MatchupScraper`, and `EventDataScraper` abstractions to be reused across all data providers (#80)
- Added `Feed` and `Provider` enums to catalog sources in code (#80)
### Changed
- Updated fox sports extraction to use new `Runner` and `Scraper` pattern (#80)
- Additional validation in fox sports extraction (#80)
- Updated baseball reference mlb extraction to use new `Runner` and `Scraper` pattern (#82)
- Updated basketball reference nba extraction to use new `Runner` and `Scraper` pattern (#82)
- Code organization improvements (#82)

## [0.10.0] - 2025-05-29
### Changed
- Update `MLBProbableStartingPitcher` fields (#77)
- Update `MLBProbableStartingPitcher` composite key from `EventID` to `EventID` and `TeamID` (#77)
### Fixed
- Skip TBD probable starting pitchers (#77)

## [0.9.0] - 2025-05-27
### Added
- `MLBMatchupComparison` for fox sports json response for probable starting pitchers (#73)
- `MLBProbableStartingPitcher` fox sports data model for probable starting pitcher (#73)
- Scraping fox sports MLB probable starting pitcher using `MLBProbableStartingPitcherScraper` (#73)
### Fixed
- Ignore baseball-reference MLB all-star matchups (#74)

## [0.8.2] - 2025-05-17
### Fixed
- Optimize extraction of matchups on baseball-reference and basketball-reference (#70)

## [0.8.1] - 2025-05-14
### Fixed
- Fixed error when scraping mlb matchup summary tables with events that completed at later dates on baseball reference (#68)

## [0.8.0] - 2025-05-14
### Fixed
- Fixed "invalid memory address or nil pointer dereference" in foxsports nba box score stats scraper (#65)
### Changed
- fox sports `NBABoxScoreStats.Position` is now optional (#65)

## [0.7.3] - 2025-04-30
### Changed
- Commit mocks generated by mockery (#62)

## [0.7.2] - 2025-04-30
### Changed
- Updated package name for `util/sportsreference/base_runner_test.go` from `sportsreferenceutil` to `sportsreferenceutil_test` (#60)

## [0.7.1] - 2025-04-22
### Added
- added `convertedtype=UTF8` to all parquet string field tags (#58)

## [0.7.0] - 2025-04-17
### Added
- Created fox sports `MLBEventData` jsonresponse (#56)
- Created fox sports `MLBPitchingBoxScoreStats` and `MLBBattingBoxScoreStats` data models (#56)
- Created `MLBBattingBoxScoreScraper` to scrape fox sports MLB batting box score data (#56)
- Created `MLBPitchingBoxScoreScraper` to scrape fox sports MLB pitching box score data (#56)

### Changed
- Generalized fox sports jsonresponse `BoxScoreStatline` and `BoxScoreStats` for mlb and nba use cases (#56)
- json omit fields ending in "Parquet" (#56)

## [0.6.0] - 2025-04-15
### Added
- Added `github.com/xitongsys/parquet-go` dependency (#52)
- Added parquet tags to data models (#52)
- Added `EventTime` field to fox sports `NBABoxScoreStats` (#52)
- Added `PullTimestampParquet` field to all data models for parquet conversion (#52)
- Data models with `EventTime` also include `EventTimeParquet` for parquet conversion (#52)
- Data models with `EventDate` also include `EventDateParquet` for parquet conversion (#52)
- Added utility function `TimeToDays` (#52)
- Created `TextToInt32` utility function (#52)

### Changed
- Changed `int` type instances in data models to `int32` to match parquet writer type expectations (#52)
- Less noise in logging for NBA basic box score stats and advanced box score stats when parsed text is empty string. (#52)
- Update to documentation to included supported file formats to export/import. (#52)

### Fixed
- Correct default assignment for `threePointAttempts` of basketball reference NBA basic box score stats (#52)

## [0.5.0] - 2025-04-10
### Added
- Created `Runner` and `Scraper` abstractions to retrieve and parse fox sports event data. (#50)
- Created `NBAEventData` JSON response payload for fox sports NBA event data. (#50)
- Created `NBABoxScoreStats` data model (#50)
- Created `NBABoxScoreScraper` interface to scrape fox sports NBA box score data (#50)

### Changed
- Renamed `selectionId` to `segmentID`, `Id` to `ID`, and `SelectionId` to `SegmentID` (#50)
- Documentation updates (#50)

## [0.4.0] - 2025-04-07
### Added
- Created fox sports matchup scraper for NBA, NFL, NCAAB, and MLB called `GeneralMatchupRunner` (#34)
- Added a JSON representation of the response payload when requesting fox sports matchup data (#34)
- Created a data model that represents fox sports matchup (#34)

### Changed
- Documentation updates (#34)

## [0.3.0] - 2025-04-04
### Changed
- Moved `tools/` to `internal/tools/` to avoid exporting (#32)

## [0.2.0] - 2025-04-03
### Added
- Created `Round` util function (#30)
- Support for scraping all periods to NBA `BasicBoxScoreRunner` (#30)
### Changed
- 2 decimal place precision for `transformMinutesPlayed` (#30)
- Documentation updates (#30)

## [0.1.0] - 2025-04-01
### Added
- Dependencies for tooling (#21):
    - github.com/go-git/go-git/v5
    - github.com/spf13/cobra
    - golang.org/x/term
    - github.com/Masterminds/semver/v3
- Semantic versioning embedded in project (#21)
### Changed
- Documentation update (#21)

## [0.1.0-beta.4] - 2025-03-20
### Changed
- Documentation updates
### Fixed
- `ERROR: unhandled page event *page.EventFrameStartedNavigating` fixed in `github.com/chromedp/chromedp@v0.13.2` (https://github.com/lightning-dabbler/sportscrape/pull/19)

## [0.1.0-beta.3] - 2025-03-19
### Changed
- Minor update to `go.mod` and `go.sum` based on introduction of mockery (https://github.com/lightning-dabbler/sportscrape/pull/16)
- Documentation updates (https://github.com/lightning-dabbler/sportscrape/pull/17)

## [0.1.0-beta.2] - 2025-03-18
### Changed
- Documentation updates and corrections

## [0.1.0-beta.1] - 2025-03-18
### Added
- Added `EventID` to sportsreference util and updated `MatchupRunner.GetMatchups(...)` to use it.
- Created `MLBMatchup` model and `MatchupRunner` to scrape MLB matchups from https://baseball-reference.com
- Created `MLBBattingBoxScoreStats` and `BattingBoxScoreRunner` to scrape MLB batting box score data from https://baseball-reference.com
- Created `MLBPitchingBoxScoreStats` model and `PitchingBoxScoreRunner` to scrape MLB pitching box score data from https://baseball-reference.com
### Changed
- Moved local `extractPlayerID` to sportsreference util and renamed it `PlayerID`
- Changed `headerValues` type from `map[string]string` to `[]string`
- renamed `headerValues` to `Headers`
- Moved `Headers` to sportsreferenceutil package

## [0.1.0-beta] - 2025-03-02
### Added
- Created `Runner`, `MatchupRunner`, and `BoxScoreRunner` in base_runner.go as sportreference utils to better configure and facilitate scraping for basketball-reference.com matchups and box scores (#9)
- Created `RetrieverOption` to make `NewDocumentRetriever` more configurable with default values to keep the function signature consistent long term (#9)
### Changed
- Updated `NewDocumentRetriever` function signature (#9)
- Refactored each scraping process to inherit from `MatchupRunner` and `BoxScoreRunner` (#9)
- Each box score stats scraper uses `Processor.GetSegmentBoxScoreStats` to orchestrate scraping. (#9)
- All round better configuration (#9)
- Update logging all over to use `log` (#9)

## [0.1.0-alpha.1] - 2025-03-02
### Added
- Added `GetAdvBoxScoreStats` to scrape NBA advanced box score stats from basketball-reference.com (#5)
- `NBAAdvBoxScoreStats` is the data model for NBA Advanced box score stats from basketball-reference.com (#5)
- Created shared.go for shared components for scraping NBA box score data from basketball-reference.com (#5)
- Created functions to extract player ID and transform player minutes played (#5)
### Changed
- Refactored `getBasicBoxScoreStats` and `GetMatchups` to use `networkHeaders` and `DocumentRetriever` (#5)
- Logging is more concise and better for `getBasicBoxScoreStats` (#5)
- Slight refactor of `getBasicBoxScoreStats` to use shared components (#5)

## [0.1.0-alpha] - 2025-02-25
### Added
- Functionality to scrape NBA matchup and related basic box score stats data from https://basketball-reference.com (#2)
- Utility functions for data transformations and requests
