package elements

// The AttachmentId element identifies an item or file attachment. This element is used in CreateAttachment responses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentid
type AttachmentId struct {
	// Identifies the unique identifier of the attachment.
	Id *string `xml:"Id,attr"`
	// Identifies the change key of the root store item to which the attachment is attached.
	RootItemChangeKey *string `xml:"RootItemChangeKey,attr"`
	// Identifies the unique identifier of the root store item to which the attachment is attached.
	RootItemId *string `xml:"RootItemId,attr"`
}
