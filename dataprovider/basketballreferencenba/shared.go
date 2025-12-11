package basketballreferencenba

import (
	"github.com/chromedp/cdproto/network"
)

var networkHeaders network.Headers = network.Headers(map[string]interface{}{
	"authority":                 "www.basketball-reference.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language":           "en-US;q=0.8",
	"cookie":                    "is_live=true; sr_note_box_countdown=57",
	"if-modified-since":         "Tue, 08 Nov 2022 01:08:31 GMT",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "none",
	"sec-fetch-user":            "?1",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
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
