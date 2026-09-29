package jsonresponse

// Teams is the response of https://api.nfl.com/experience/v1/teams?season={season}
type Teams struct {
	Teams []TeamInfo `json:"teams"`
}

type TeamInfo struct {
	ID           string `json:"id"`
	Abbreviation string `json:"abbreviation"`
	FullName     string `json:"fullName"`
	NickName     string `json:"nickName"`
}
