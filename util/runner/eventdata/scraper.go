package eventdata

type Scraper interface {
	Scrape(matchup interface{}) OutputWrapper
	Init()
}
