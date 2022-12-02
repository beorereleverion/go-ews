package elements

// The MarkAllItemsAsReadResponse element specifies the response to a MarkAllItemsAsRead request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markallitemsasreadresponse
type MarkAllItemsAsReadResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
