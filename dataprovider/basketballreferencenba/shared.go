package basketballreferencenba

import (
	"github.com/chromedp/cdproto/network"
)

var NetworkHeaders network.Headers = network.Headers(map[string]any{
	"authority":                 "www.basketball-reference.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
	"accept-language":           "en-US;q=0.7",
	"sec-ch-ua-mobile":          "?0",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
})

const (
	// https://www.basketball-reference.com/boxscores/{event_id}.html
	contentReadySelector         = "#content"
	boxScoreStatsRecordsSelector = `tbody > tr`
	boxScoreStarterHeaders       = `thead > tr:nth-child(2) th`
	boxScoreReserveHeaders       = `th`
	boxScorePlayerSelector       = "th"
	boxScorePlayerLinkSelector   = boxScorePlayerSelector + " > a"
)
