package elements

// The HasAttachment element specifies a Boolean value to indicate whether the item has attachments.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasattachment
type HasAttachment bool

const (
	// false
	HasAttachmentfalse HasAttachment = false
	// true
	HasAttachmenttrue HasAttachment = true
)
