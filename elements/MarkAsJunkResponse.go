package elements

// The MarkAsJunkResponse element specifies the response to a MarkAsJunk request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasjunkresponse
type MarkAsJunkResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
