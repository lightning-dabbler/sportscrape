package jsonresponse

// Person is the response of https://api.nfl.com/football/v2/persons/{person_id}
type Person struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	GSISID      string `json:"gsisId"`
}
