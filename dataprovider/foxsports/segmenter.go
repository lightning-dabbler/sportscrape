package foxsports

import (
	"time"
)

// Segmenter is the interface for constructing Segment IDs
type Segmenter interface {
	// GetSegmentID returns the ID that is concatenated to a League's URL to fetch the relevant point-in-time dataset.
	GetSegmentID() (string, error)
}

type GeneralSegmenter struct {
	Date string
}

func (cs *GeneralSegmenter) GetSegmentID() (string, error) {
	date, err := time.Parse(time.DateOnly, cs.Date)
	if err != nil {
		return "", err
	}
	return date.Format("20060102"), nil
}

type NFLSegmenter struct {
	Season SeasonType
	Year   int32
	Week   int32
}

func (nfls *NFLSegmenter) GetSegmentID() (string, error) {
	return nfls.Season.SegmentID(nfls.Year, nfls.Week), nil
}
