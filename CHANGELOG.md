# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0-beta.3] - 2025-03-19
### Changed
- Minor update to `go.mod` and `go.sum` based on introduction of mockery
- Documentation updates

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
