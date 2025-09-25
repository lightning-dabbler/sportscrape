package jsonresponse

import (
	"encoding/json"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
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
	PullTime time.Time       `json:"pull_time"`
	Raw      json.RawMessage `json:"-"`
	League   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"league"`

	Events map[string][]ESPNMMAEvent `json:"events"`
}

// filter events where link is not empty
func (f *ESPNMMASchedule) FilterScrapeableEvents() []ESPNMMAEvent {
	var events []ESPNMMAEvent
	for _, event := range f.Events {
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

	for _, event := range f.Events {
		for _, e := range event {
			if e.Link != "" {
				eventTime, _ := time.Parse("2006-01-02T15:04Z", e.Date)

				m := model.Matchup{
					PullTimestamp:        f.PullTime,
					PullTimestampParquet: f.PullTime.UnixNano() / 1000000,
					EventID:              e.ID,
					EventTime:            eventTime,
					EventTimeParquet:     eventTime.UnixNano() / 1000000,
					LeagueID:             f.League.ID,
					LeagueName:           f.League.Name,
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
