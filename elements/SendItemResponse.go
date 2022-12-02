package elements

// The SendItemResponse element defines a response to a SendItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senditemresponse
type SendItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
