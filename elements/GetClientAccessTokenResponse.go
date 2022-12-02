package elements

// The GetClientAccessTokenResponse element contains the response to a GetClientAccessToken operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstokenresponse
type GetClientAccessTokenResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
