package elements

// The GetMailTipsResponse element represents the response message for a GetMailTips operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getmailtipsresponse
type GetMailTipsResponse struct {
	// The ResponseMessages element represents a list of mail tips response messages.
	ResponseMessages *ResponseMessagesArrayOfMailTipsResponseMessageType `xml:"m:ResponseMessages"`
}
