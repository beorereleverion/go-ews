package elements

// The GetEventsResponse element represents a response to a GetEvents request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/geteventsresponse
type GetEventsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
