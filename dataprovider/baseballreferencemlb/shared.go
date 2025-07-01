package baseballreferencemlb

import (
	"fmt"
	"regexp"

	"github.com/chromedp/cdproto/network"
)

var networkHeaders network.Headers = network.Headers(map[string]interface{}{
	"authority":                 "www.baseball-reference.com",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
	"accept-language":           "en-US,en;q=0.9",
	"cache-control":             "max-age=0",
	"cookie":                    "is_live=true; sr_note_box_countdown=64",
	"if-modified-since":         "Sun, 19 Mar 2023 21:17:33 GMT",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
})

// StatType enum for different types of MLB statistics to scrape
type StatType int

const (
	Batting StatType = iota
	Pitching
)

const (
	contentReadySelector = "#content"
	headersSelector      = "thead > tr th"
	recordSelector       = "tbody > tr"
	positionSelector     = "th"
	playerSelector       = positionSelector + " a"
)

func (st StatType) String() string {
	switch st {
	case Batting:
		return "batting"
	case Pitching:
		return "pitching"
	default:
		return "unknown"
	}
}

var replaceStringRegex = regexp.MustCompile(`\.|\s`)

// generateStatTableSelector generates the stat table id selector used to select a teams's stat table

// Parameter:
//   - team: The name of the MLB team given by baseball-reference.com
//   - statType: The type of stat to generate an id selector for
//
// Returns a string representing the id selector
func generateStatTableSelector(team string, statType StatType) string {
	return fmt.Sprintf("#%s%s", replaceStringRegex.ReplaceAllString(team, ""), statType.String())
}
