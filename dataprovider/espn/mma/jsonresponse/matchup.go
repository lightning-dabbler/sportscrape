package jsonresponse

import (
	"encoding/json"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/xitongsys/parquet-go/types"
)

type ESPNMMAEvent struct {
	ID                     string `json:"id"`
	Date                   string `json:"date"`
	Completed              bool   `json:"completed"`
	Link                   string `json:"link"`
	Name                   string `json:"name"`
	IsPostponedOrCancelled bool   `json:"isPostponedOrCanceled"`
	Status                 struct {
		ID     string `json:"id"`
		State  string `json:"state"`
		Detail string `json:"detail"`
	}
}

type ESPNMMASchedule struct {
	PullTime time.Time       `json:"-"`
	Raw      json.RawMessage `json:"-"`
	Page     struct {
		Content struct {
			League struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"league"`
			Events map[string][]ESPNMMAEvent `json:"events"`
		} `json:"content"`
	} `json:"page"`
}

// filter events where link is not empty
func (f *ESPNMMASchedule) FilterScrapeableEvents() []ESPNMMAEvent {
	var events []ESPNMMAEvent
	for _, event := range f.Page.Content.Events {
		for _, e := range event {
			if e.Link != "" {
				events = append(events, e)
			}
		}
	}
	return events
}

func (f *ESPNMMASchedule) GetScrapableMatchup() []model.Matchup {
	var events []model.Matchup

	for _, event := range f.Page.Content.Events {
		for _, e := range event {
			if e.Link != "" {
				eventTime, _ := time.Parse("2006-01-02T15:04Z", e.Date)

				m := model.Matchup{
					PullTimestamp:        f.PullTime,
					PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(f.PullTime, true),
					EventID:              e.ID,
					EventTime:            eventTime,
					EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTime, true),
					LeagueID:             f.Page.Content.League.ID,
					LeagueName:           f.Page.Content.League.Name,
					//Date: f,
					Completed:              e.Completed,
					Link:                   e.Link,
					Name:                   e.Name,
					IsPostponedOrCancelled: e.IsPostponedOrCancelled,
					StatusID:               e.Status.ID,
					StatusState:            e.Status.State,
					StatusDetail:           e.Status.Detail,
				}
				events = append(events, m)
			}
		}
	}
	return events
}
