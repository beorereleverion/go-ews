package elements

// The CreateItemResponse element defines a response to a CreateItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitemresponse
type CreateItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
