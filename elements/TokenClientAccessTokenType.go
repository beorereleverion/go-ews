package elements

// The Token element specifies a client access token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/token-clientaccesstokentype
type TokenClientAccessTokenType struct {
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"m:ID"`
	// The TTL element specifies the time, in minutes, that the token is valid.
	TTL *TTL `xml:"t:TTL"`
	// The TokenType element specifies the type of token.
	TokenType *TokenType `xml:"t:TokenType"`
	// The TokenValue element specifies the encoded client access token.
	TokenValue *TokenValue `xml:"t:TokenValue"`
}
