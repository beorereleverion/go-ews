package elements

// The MoveItemResponse element defines a response to a MoveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moveitemresponse
type MoveItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
