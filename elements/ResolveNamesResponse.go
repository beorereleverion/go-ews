package elements

// The ResolveNamesResponse element defines a response to a ResolveNames request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolvenamesresponse
type ResolveNamesResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
