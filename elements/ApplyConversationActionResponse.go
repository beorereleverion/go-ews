package elements

// The ApplyConversationActionResponse element defines a response to an ApplyConversationAction operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applyconversationactionresponse
type ApplyConversationActionResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
