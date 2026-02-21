//go:build unit

package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test model with a Raw []byte field to verify setBytesField is applied by HydrateModel.
type testModelWithRaw struct {
	Name string `json:"name"`
	Raw  []byte
}

// Test model without a Raw field to verify HydrateModel still succeeds.
type testModelNoRaw struct {
	Name string `json:"name"`
}

func TestRetrieveBytes_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	s := BaseJsonScraper[testModelNoRaw]{}
	got, err := s.RetrieveBytes(ts.URL)
	require.NoError(t, err, "RetrieveBytes() unexpected error")
	require.NotNil(t, got, "RetrieveBytes() returned nil slice")
	assert.Equal(t, []byte("hello world"), *got, "RetrieveBytes() mismatch")
}

func TestRetrieveModel_Success(t *testing.T) {
	const payload = `{"name":"alpha"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(payload))
	}))
	defer ts.Close()

	s := BaseJsonScraper[testModelWithRaw]{}
	got, err := s.RetrieveModel(ts.URL)
	require.NoError(t, err, "RetrieveModel() unexpected error")
	require.NotNil(t, got, "RetrieveModel() returned nil model")
	assert.Equal(t, "alpha", got.Name)
	assert.Equal(t, []byte(payload), got.Raw)
}

func TestHydrateModel_WithRawField_SetsRaw(t *testing.T) {
	const payload = `{"name":"beta"}`
	s := BaseJsonScraper[testModelWithRaw]{}

	model, err := s.HydrateModel([]byte(payload))
	require.NoError(t, err, "HydrateModel() unexpected error")
	assert.Equal(t, "beta", model.Name)
	assert.Equal(t, []byte(payload), model.Raw)
}

func TestHydrateModel_WithoutRawField_IgnoresMissingField(t *testing.T) {
	const payload = `{"name":"gamma"}`
	s := BaseJsonScraper[testModelNoRaw]{}

	model, err := s.HydrateModel([]byte(payload))
	require.NoError(t, err, "HydrateModel() unexpected error")
	assert.Equal(t, "gamma", model.Name)
}

func TestHydrateModel_InvalidJSON_ReturnsError(t *testing.T) {
	const payload = `{"name":` // invalid JSON
	s := BaseJsonScraper[testModelWithRaw]{}

	model, err := s.HydrateModel([]byte(payload))
	assert.Error(t, err, "HydrateModel() expected error")
	assert.Nil(t, model)
}

// Additional direct tests for setBytesField to cover edge cases.

func TestSetBytesField_SetsFieldSuccessfully(t *testing.T) {
	type s struct {
		Raw []byte
	}
	instance := &s{}
	data := []byte("data")
	err := setBytesField(instance, "Raw", data)
	require.NoError(t, err, "setBytesField() unexpected error")
	assert.Equal(t, data, instance.Raw)
}

func TestSetBytesField_Errors(t *testing.T) {
	type withOther struct {
		Other []byte
	}
	type wrongType struct {
		Raw string
	}
	type unexported struct {
		raw []byte
	}

	tests := []struct {
		name      string
		target    any
		field     string
		expectErr string
	}{
		{
			name:      "non-pointer input",
			target:    struct{ Raw []byte }{},
			field:     "Raw",
			expectErr: "expected a non-nil pointer to a struct",
		},
		{
			name:      "nil pointer",
			target:    (*struct{ Raw []byte })(nil),
			field:     "Raw",
			expectErr: "expected a non-nil pointer to a struct",
		},
		{
			name:      "pointer to non-struct",
			target:    new(int),
			field:     "Raw",
			expectErr: "expected a pointer to a struct",
		},
		{
			name:      "missing field",
			target:    &withOther{},
			field:     "Raw",
			expectErr: "field 'Raw' not found",
		},
		{
			name:      "unexported field not settable",
			target:    &unexported{},
			field:     "raw",
			expectErr: "field 'raw' is not settable",
		},
		{
			name:      "wrong field type",
			target:    &wrongType{},
			field:     "Raw",
			expectErr: "is not of type []byte",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := setBytesField(tt.target, tt.field, []byte("x"))
			require.Error(t, err, "setBytesField() expected error")
			assert.ErrorContains(t, err, tt.expectErr)
		})
	}
}
