package elements

// The FindItemResponse element defines a response to a FindItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditemresponse
type FindItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
