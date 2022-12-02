package elements

// The SubscribeResponse element defines a response to a Subscribe request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriberesponse
type SubscribeResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
