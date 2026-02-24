package feed

import (
	"context"
	"fmt"
)

type ProviderExtractor interface {
	Scrape(ctx context.Context) error
	ValidateFeed() error
}

var ErrUnsupportedFeed error = fmt.Errorf("unsupported data feed")
