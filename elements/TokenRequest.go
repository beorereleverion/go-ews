package elements

// The TokenRequest element specifies a single token request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokenrequest
type TokenRequest struct {
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"m:ID"`
	// The TokenType element identifies the type of client access token that will be returned in the GetClientAccessToken response.
	TokenType *TokenTypeClientAccessTokenType `xml:"t:TokenType"`
}
