package elements

// The GetAttachmentResponse element defines a response to a GetAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachmentresponse
type GetAttachmentResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
