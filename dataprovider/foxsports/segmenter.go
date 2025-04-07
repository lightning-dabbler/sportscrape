package foxsports

import (
	"time"
)

type GeneralSementer struct {
	Date string
}

func (cs *GeneralSementer) GetSegmentId() (string, error) {
	date, err := time.Parse("2006-01-02", cs.Date)
	if err != nil {
		return "", err
	}
	return date.Format("20060102"), nil
}

type NFLSementer struct {
	Season SeasonType
	Year   int32
	Week   int32
}

func (nfls *NFLSementer) GetSegmentId() (string, error) {
	return nfls.Season.SelectionId(nfls.Year, nfls.Week), nil
}
