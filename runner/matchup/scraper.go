package matchup

type Scraper interface {
	Scrape() OutputWrapper
	Init()
}
