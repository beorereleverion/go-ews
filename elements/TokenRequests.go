package elements

// The TokenRequests element contains an array of token requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokenrequests
type TokenRequests struct {
	// The TokenRequest element specifies a single token request.
	TokenRequest *TokenRequest `xml:"t:TokenRequest"`
}
