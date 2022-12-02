package elements

// The CreateAttachmentResponse element defines a response to a CreateAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createattachmentresponse
type CreateAttachmentResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
