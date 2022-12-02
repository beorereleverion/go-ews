package elements

// The GetClientAccessToken element contains a request to get a client access token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstoken
type GetClientAccessToken struct {
	// The TokenRequests element contains an array of token requests.
	TokenRequests *TokenRequests `xml:"m:TokenRequests"`
}
