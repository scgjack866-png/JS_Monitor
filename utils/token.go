package utils

type Token struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
	Expires      string `json:"expires"`
}
