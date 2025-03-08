package mlb

import (
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
