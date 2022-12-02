package elements

// The DeleteItemResponse element defines a response to a single DeleteItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemresponse
type DeleteItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
