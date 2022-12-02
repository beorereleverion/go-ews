package elements

// The GetConversationItemsResponse element defines a response to a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getconversationitemsresponse
type GetConversationItemsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
