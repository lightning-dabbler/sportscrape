package jsonresponse

// Token is the response of https://api.nfl.com/identity/v3/token
type Token struct {
	AccessToken string `json:"accessToken"`
	// ExpiresIn is the token's expiry in epoch seconds
	ExpiresIn int64 `json:"expiresIn"`
}
