package elements

// The GetItemResponse element defines a response to a GetItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitemresponse
type GetItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
