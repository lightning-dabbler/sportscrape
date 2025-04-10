package foxsports

import (
	"time"
)

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
