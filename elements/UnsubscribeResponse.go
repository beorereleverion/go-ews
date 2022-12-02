package elements

// The UnsubscribeResponse element defines a response to an Unsubscribe request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unsubscriberesponse
type UnsubscribeResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
