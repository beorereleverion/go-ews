package elements

// The DeleteAttachment element is the root element in a request to delete an attachment from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteattachment
type DeleteAttachment struct {
	// The AttachmentIds element contains an array of attachment identifiers.
	AttachmentIds *AttachmentIds `xml:"m:AttachmentIds"`
}
