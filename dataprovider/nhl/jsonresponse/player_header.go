package jsonresponse

// PlayerHeader is the json response from https://api-web.nhle.com/v2/player/{player_id}/header
type PlayerHeader struct {
	PlayerID  int64           `json:"playerId"`
	FirstName LocalizedString `json:"firstName"`
	LastName  LocalizedString `json:"lastName"`
}
