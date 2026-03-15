package request

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// DocumentRetrieverV2 manages a persistent headless Chrome browser session for
// fetching and parsing web documents. Unlike DocumentRetriever, the browser
// process and cookies are shared across all RetrieveDocument calls.
// Each RetrieveDocument call opens a new tab within the shared browser session,
// allowing safe concurrent use. Create with NewDocumentRetrieverV2 and call
// Close when done.
type DocumentRetrieverV2 struct {
	Timeout        time.Duration
	Debug          bool
	ChromeRun      func(ctx context.Context, actions ...chromedp.Action) error
	DocumentReader func(r io.Reader) (*goquery.Document, error)
	NewTabContext  func(ctx context.Context) (context.Context, context.CancelFunc)
	// Persistent browser context — parent for all per-call tab contexts
	browserCtx     context.Context
	browserCancel  context.CancelFunc
	networkHeaders network.Headers
}

// RetrieverOptionV2 is a functional option for configuring a DocumentRetrieverV2.
type RetrieverOptionV2 func(*DocumentRetrieverV2)

// WithTimeoutV2 returns a RetrieverOptionV2 that sets the timeout duration for document retrieval operations.
//
// Parameter:
//   - timeout: The duration after which document retrieval operations will be cancelled
func WithTimeoutV2(timeout time.Duration) RetrieverOptionV2 {
	return func(dr *DocumentRetrieverV2) {
		dr.Timeout = timeout
	}
}

// WithDebugV2 returns a RetrieverOptionV2 that enables or disables debug logging.
//
// Parameter:
//   - debug: When true, enables verbose logging of Chrome operations
func WithDebugV2(debug bool) RetrieverOptionV2 {
	return func(dr *DocumentRetrieverV2) {
		dr.Debug = debug
	}
}

// NewDocumentRetrieverV2 creates and initializes a persistent browser session,
// storing network headers for per-tab setup and applying any provided options.
// Call Close() when done.
func NewDocumentRetrieverV2(networkHeaders network.Headers, options ...RetrieverOptionV2) (*DocumentRetrieverV2, error) {
	dr := &DocumentRetrieverV2{
		Timeout:        1 * time.Minute,
		Debug:          false,
		ChromeRun:      chromedp.Run,
		DocumentReader: goquery.NewDocumentFromReader,
		NewTabContext: func(ctx context.Context) (context.Context, context.CancelFunc) {
			return chromedp.NewContext(ctx)
		},
		networkHeaders: networkHeaders,
	}
	for _, option := range options {
		option(dr)
	}

	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), allocOpts...)

	var contextOptions []chromedp.ContextOption
	if dr.Debug {
		contextOptions = append(contextOptions, chromedp.WithDebugf(log.Printf))
	}

	browserCtx, browserCancel := chromedp.NewContext(allocCtx, contextOptions...)

	// Warm up the browser process with an initial target so subsequent tab
	// creation is fast.
	if err := chromedp.Run(browserCtx); err != nil {
		allocCancel()
		browserCancel()
		return nil, fmt.Errorf("failed to initialize browser session: %w", err)
	}

	dr.browserCtx = browserCtx
	dr.browserCancel = func() {
		browserCancel()
		allocCancel()
	}

	return dr, nil
}

// Close tears down the persistent browser session.
func (dr *DocumentRetrieverV2) Close() {
	if dr.browserCancel != nil {
		dr.browserCancel()
	}
}

// RetrieveDocument opens a new tab within the existing browser session,
// navigates to url, waits for waitReadySelector, and returns the parsed
// document. The tab is closed when the call returns. Safe to call concurrently.
func (dr *DocumentRetrieverV2) RetrieveDocument(url string, waitReadySelector string) (*goquery.Document, error) {
	tabCtx, tabCancel := dr.NewTabContext(dr.browserCtx)
	defer tabCancel()

	ctx, cancel := context.WithTimeout(tabCtx, dr.Timeout)
	defer cancel()

	slog.Info("Retrieving document", "url", url)
	var outer string
	if err := dr.ChromeRun(ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(dr.networkHeaders),
		chromedp.Navigate(url),
		chromedp.WaitReady(waitReadySelector),
		chromedp.OuterHTML(waitReadySelector, &outer, chromedp.ByQuery),
	); err != nil {
		return nil, fmt.Errorf("error fetching document from %s: %w", url, err)
	}

	doc, err := dr.DocumentReader(strings.NewReader(outer))
	if err != nil {
		return nil, fmt.Errorf("error reading document from %s: %w", url, err)
	}
	return doc, nil
}
