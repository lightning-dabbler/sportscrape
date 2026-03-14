//go:build unit

package request

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/stretchr/testify/assert"
)

// TestWithTimeoutV2Option verifies that WithTimeoutV2 sets Timeout on the retriever.
func TestWithTimeoutV2Option(t *testing.T) {
	timeouts := []time.Duration{
		500 * time.Millisecond,
		1 * time.Second,
		5 * time.Minute,
	}

	for _, timeout := range timeouts {
		dr := &DocumentRetrieverV2{}
		option := WithTimeoutV2(timeout)
		option(dr)
		assert.Equal(t, timeout, dr.Timeout, "Timeout should match the value passed to WithTimeoutV2")
	}
}

// TestWithDebugV2Option verifies that WithDebugV2 sets Debug on the retriever.
func TestWithDebugV2Option(t *testing.T) {
	testCases := []struct {
		name  string
		debug bool
	}{
		{"Enable Debug", true},
		{"Disable Debug", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dr := &DocumentRetrieverV2{}
			option := WithDebugV2(tc.debug)
			option(dr)
			assert.Equal(t, tc.debug, dr.Debug, "Debug should match the value passed to WithDebugV2")
		})
	}
}

// TestDocumentRetrieverV2RetrieveDocument tests RetrieveDocument with injected
// ChromeRun and DocumentReader stubs so no real browser is required.
func TestDocumentRetrieverV2RetrieveDocument(t *testing.T) {
	testCases := []struct {
		name         string
		runErr       error
		docReaderErr error
		expectErr    bool
	}{
		{
			name:      "Success",
			expectErr: false,
		},
		{
			name:      "ChromeRun error",
			runErr:    errors.New("chrome error"),
			expectErr: true,
		},
		{
			name:         "DocumentReader error",
			docReaderErr: errors.New("doc reader error"),
			expectErr:    true,
		},
		{
			name:      "Context deadline exceeded",
			runErr:    context.DeadlineExceeded,
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			dr := &DocumentRetrieverV2{
				Timeout: 100 * time.Millisecond,
				ChromeRun: func(ctx context.Context, actions ...chromedp.Action) error {
					return tt.runErr
				},
				DocumentReader: func(r io.Reader) (*goquery.Document, error) {
					if tt.docReaderErr != nil {
						return nil, tt.docReaderErr
					}
					return goquery.NewDocumentFromReader(strings.NewReader("<html><body>Test</body></html>"))
				},
				NewTabContext: func(parent context.Context) (context.Context, context.CancelFunc) {
					return context.WithCancel(parent)
				},
				browserCtx: ctx,
			}

			doc, err := dr.RetrieveDocument("https://example.com", "body")

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, doc)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, doc)
			}
		})
	}
}

// TestDocumentRetrieverV2Close verifies that Close invokes browserCancel exactly
// once and is safe to call when browserCancel is nil.
func TestDocumentRetrieverV2Close(t *testing.T) {
	t.Run("nil browserCancel is a no-op", func(t *testing.T) {
		dr := &DocumentRetrieverV2{}
		assert.NotPanics(t, func() { dr.Close() })
	})

	t.Run("non-nil browserCancel is called", func(t *testing.T) {
		called := false
		dr := &DocumentRetrieverV2{
			browserCancel: func() { called = true },
		}
		dr.Close()
		assert.True(t, called, "browserCancel should have been invoked by Close")
	})
}
