# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
- Refactored `getBasicBoxScoreStats` and `GetMatchups` to use `networkHeaders` and `DocumentRetriever`
- Logging is more concise for `getBasicBoxScoreStats`
- Slight refactor of `getBasicBoxScoreStats` to use shared components
### Added
- Added `GetAdvBoxScoreStats` to scrape NBA advanced box score stats from basketball-reference.com
- `NBAAdvBoxScoreStats` is the data model for NBA Advanced box score stats from basketball-reference.com
- Created shared.go for shared components for scraping NBA box score data from basketball-reference.com
- Created functions to extract player ID and transform player minutes played

## [0.1.0-alpha] - 2025-02-25
### Added
- Functionality to scrape NBA matchup and related basic box score stats data from https://basketball-reference.com (#2)
- Utility functions for data transformations and requests
