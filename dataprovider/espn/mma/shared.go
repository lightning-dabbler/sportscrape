package mma

import "github.com/chromedp/cdproto/network"

var NetworkHeaders network.Headers = network.Headers(map[string]any{
	"authority":       "www.espn.com",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
	"accept-language": "en-US;q=0.8",
	// "cookie":                    "is_live=true; sr_note_box_countdown=57",
	// "if-modified-since":         "Tue, 08 Nov 2022 01:08:31 GMT",
	// "sec-fetch-dest":            "document",
	// "sec-fetch-mode":            "navigate",
	// "sec-fetch-site":            "none",
	// "sec-fetch-user":            "?1",
	"sec-ch-ua-mobile":          "?0",
	"sec-gpc":                   "1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",
})
