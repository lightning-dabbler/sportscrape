# CLAUDE.md

Guidance for AI coding assistants working in this repository.

## What this repo is

Go library and CLI (`github.com/lightning-dabbler/sportscrape`) for scraping
sports data. Providers live in `dataprovider/<provider>/`. The CLI lives in
`cmd/sportscrape/`. Root-level `catalog.go`, `context.go`, and `output.go`
define shared enumerations and types used across the library.

---

## Commands

```bash
# Build
make build-sportscrape       # → bin/sportscrape
make build-tools             # → bin/tools

# Test — plain `go test ./...` produces NO output (build tags required)
make unit-tests              # -tags=unit, with coverage
make all-tests               # -tags="unit integration", with coverage

# Generate
make mocks-gen               # regenerate internal/mocks/ via mockery
```

---

## Non-obvious rules

### Build tags
- **`go test ./...` with no tags matches zero files and silently produces nothing.**
- Unit test files: `//go:build unit` on line 1, before `package`
- Integration test files (also included in unit suite): `//go:build unit || integration` on line 1
- The build tag must be the absolute first line of the file.

### Scraper lifecycle — Init/Close ownership
- **Constructors must not call `Init()`.**
- `Init()` and `Close()` are the caller's responsibility — or the runner's (see below).
- `BaseDocumentScraper.Init()` calls `log.Fatalln` if `Timeout == 0` and no
  `DocumentRetriever` is injected. Always set `Timeout` before calling `Init()`.
- `BaseDocumentScraper.Close()` sets `DocumentRetriever` to nil; it is safe to
  call when already nil.

### Runner owns Init/Close for EventDataScraper
- `EventDataRunner.Run()` calls `Init()` then `defer Close()` automatically unless
  `EventDataRunnerConfig.KeepAlive == true`, in which case the caller must call `Close()` manually.
- `MatchupRunner.Run()` calls `Init()` then `defer Close()` automatically unless
  `MatchupRunnerConfig.KeepAlive == true`, in which case the caller must call `Close()` manually.
- **Do not call `Init()` before passing a scraper to a runner.**
- `EventDataRunner` defaults `Concurrency` to `1` — it does **not** auto-detect
  CPU count. Set `Concurrency` explicitly if parallelism is needed.

### Deprecation
- `Provider.Deprecated()` and `Feed.Deprecated()` are defined in `catalog.go`.
- Runners check `Deprecated()` at the start of `Run()` and return
  `Feed.Deprecation()` as an error if true — no scraping occurs.
- `BaseballReference` provider is deprecated. `ESPNPFLMatchups` and
  `ESPNPFLFightDetails` feeds are deprecated.
- When adding a new deprecated feed/provider, add a case to the relevant
  `Deprecated()` switch in `catalog.go`.

### Scraper interfaces return output structs, not raw slices
```go
// scraper/scraper.go
type MatchupScraper[M any] interface {
    Init()
    Close()
    Feed() sportscrape.Feed
    Provider() sportscrape.Provider
    Scrape() sportscrape.MatchupOutput[M]      // not ([]M, error)
}

type EventDataScraper[M, E any] interface {
    Init()
    Close()
    Feed() sportscrape.Feed
    Provider() sportscrape.Provider
    Scrape(matchup M) sportscrape.EventDataOutput[E]   // not ([]E, error)
}
```

`MatchupOutput[M]` carries `{Error error, Output []M, Context MatchupContext}`.
`EventDataOutput[E]` carries `{Error error, Output []E, Context EventDataContext}`.
`EventDataContext` includes `EventID`, `URL`, `AwayTeam`, `HomeTeam`, timestamps.

When implementing `Scrape()` on a `MatchupScraper`, use `MatchupContext` fields
as follows — the runner treats them differently:
- `Error` — fatal; runner returns immediately with this error
- `Context.Errors` — count of individual events that failed; runner logs and continues
- `Context.Skips` — count of events intentionally skipped; runner logs a warning only

### Registering a new Provider or Feed
Follow the pattern in `catalog.go` exactly:
```go
MyProvider Provider = "my provider"
MyProviderFeed Feed = Feed(string(MyProvider) + " feed name")
```
Add `Deprecated()` cases if applicable. Then wire the CLI subcommand in
`cmd/sportscrape/internal/cli/` and feed handler in
`cmd/sportscrape/internal/feed/`.

### Pointer receivers on chromedp-based scrapers
All scrapers that embed `BaseDocumentScraper` or `BaseScraper` must use pointer
receivers on all methods.

### Mocks
Never write mocks by hand. Run `make mocks-gen`. Mock files live in
`internal/mocks/scraper/` and are generated from `scraper.MatchupScraper` and
`scraper.EventDataScraper` per `.mockery.yaml`.

### Chrome in tests
`request.NewDocumentRetriever` and `request.NewDocumentRetrieverV2` launch a
real Chrome process and **cannot be called in unit tests**.
- Inside `package request`: inject stubs via the `ChromeRun`, `DocumentReader`,
  and `NewTabContext` fields on `DocumentRetrieverV2`. Tests that call
  `RetrieveDocument` must stub `NewTabContext` as
  `func(parent context.Context) (context.Context, context.CancelFunc) { return context.WithCancel(parent) }`
  to avoid launching a real browser.
- Outside `package request`: test nil-`DocumentRetriever` error paths and
  struct-level assertions only. Use mocks from `internal/mocks/scraper/` for
  interface-level tests.

---

## Choosing a base scraper

| Situation | Use |
|---|---|
| JSON REST API | `BaseJSONScraper` (`scraper/base_json_scraper.go`) |
| HTML, no session needed | `BaseScraper` (`scraper/base_scraper.go`) |
| HTML, session/cookies required (nba.com, espn.com, basketball-reference.com) | `BaseDocumentScraper` (`scraper/base_document_scraper.go`) |

---

## Adding a new data provider

1. Create `dataprovider/<provider>/` with: `*.go`, `model/`, optionally
   `jsonresponse/`, `example_test.go`.
2. Choose the correct base scraper (table above).
3. Add `Provider` and `Feed` constants to `catalog.go`.
4. Wire CLI in `cmd/sportscrape/internal/cli/` and feed handler in
   `cmd/sportscrape/internal/feed/`.
5. Tag all test files correctly (see Build tags above).
6. If the provider requires session headers, add a package-level
   `NetworkHeaders` variable of type `network.Headers` (imported from
   `github.com/chromedp/cdproto/network`). See `dataprovider/nba` or
   `dataprovider/espn/mma` for the pattern.
