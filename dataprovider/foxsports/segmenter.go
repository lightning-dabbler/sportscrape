package foxsports

import (
	"time"
)

type GeneralSegmenter struct {
	Date string
}

func (cs *GeneralSegmenter) GetSegmentId() (string, error) {
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

func (nfls *NFLSegmenter) GetSegmentId() (string, error) {
	return nfls.Season.SelectionId(nfls.Year, nfls.Week), nil
}
