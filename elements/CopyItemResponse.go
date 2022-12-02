package elements

// The CopyItemResponse element defines a response to a CopyItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyitemresponse
type CopyItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
