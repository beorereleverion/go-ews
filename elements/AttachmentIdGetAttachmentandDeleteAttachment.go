package elements

// The AttachmentId element identifies a single attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentid-getattachment-and-deleteattachment
type AttachmentIdGetAttachmentandDeleteAttachment struct {
	// Specifies the attachment identifier.
	Id *string `xml:"Id,attr"`
}
