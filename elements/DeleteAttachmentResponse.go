package elements

// The DeleteAttachmentResponse defines a response to a DeleteAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteattachmentresponse
type DeleteAttachmentResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
