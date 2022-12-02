package elements

// The AttachmentIds element contains an array of attachment identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentids
type AttachmentIds struct {
	// The AttachmentId element identifies a single attachment.
	AttachmentId *AttachmentIdGetAttachmentandDeleteAttachment `xml:"t:AttachmentId"`
}
