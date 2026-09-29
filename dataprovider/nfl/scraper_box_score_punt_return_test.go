//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestPuntReturnBoxScoreScraper(t *testing.T) {
	s := NewPuntReturnBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 2, len(records), "2 punt returners")
	// a fair catch without a return
	r := findPlayer(t, records, func(r model.PuntReturnBoxScore) string { return r.PlayerID }, "00-0037801")
	assert.Equal(t, "KaVontae Turpin", r.Player)
	assert.Equal(t, "K.Turpin", r.PlayerShortName)
	assert.Equal(t, int32(0), r.Returns)
	assert.Equal(t, int32(1), r.FairCatches)
}
