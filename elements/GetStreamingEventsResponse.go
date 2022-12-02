package elements

// The GetStreamingEventsResponse element represents a response to a GetStreamingEvents element request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getstreamingeventsresponse
type GetStreamingEventsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
