//go:build unit || integration

package request

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnitTests(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response string
		isError  bool
	}{
		{
			name:     "successful response",
			status:   http.StatusOK,
			response: "success response",
			isError:  false,
		},
		{
			name:     "not found response",
			status:   http.StatusNotFound,
			response: "not found",
			isError:  true,
		},
		{
			name:     "server error",
			status:   http.StatusInternalServerError,
			response: "server error",
			isError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				w.Write([]byte(tt.response))
			}))
			defer dummyServer.Close()
			resp, err := Get(dummyServer.URL)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				body, _ := io.ReadAll(resp.Body)
				resp.Body.Close()
				assert.Equal(t, tt.response, string(body))
			}
		})
	}
}

func TestGetIntegrationTests(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	tests := []struct {
		name    string
		url     string
		status  int
		isError bool
	}{
		{
			name:    "valid url",
			url:     "https://example.com",
			status:  http.StatusOK,
			isError: false,
		},
		{
			name:    "invalid protocol",
			url:     "efkjfnekfn://example.com",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := Get(tt.url)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.status, resp.StatusCode)
			}
		})
	}
}
