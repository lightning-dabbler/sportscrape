package model

import "encoding/json"

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

type ESPNMMAFeed struct {
	Raw    json.RawMessage `json:"-"`
	League struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"league"`

	Events map[string][]ESPNMMAEvent `json:"events"`
}

// filter events where link is not empty
func (f *ESPNMMAFeed) FilterScrapeableEvents() []ESPNMMAEvent {
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
