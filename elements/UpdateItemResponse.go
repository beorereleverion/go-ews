package elements

// The UpdateItemResponse element defines a response to an UpdateItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitemresponse
type UpdateItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
