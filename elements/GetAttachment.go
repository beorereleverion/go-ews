package elements

// The GetAttachment element is the root element in a request to get an attachment from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachment
type GetAttachment struct {
	// The AttachmentIds element contains an array of attachment identifiers.
	AttachmentIds *AttachmentIds `xml:"m:AttachmentIds"`
	// The AttachmentShape element identifies additional properties to return in a response to a GetAttachment request.
	AttachmentShape *AttachmentShape `xml:"m:AttachmentShape"`
	// The GetAttachment element is the root element in a request to get an attachment from the Exchange store.
	GetAttachment *GetAttachment `xml:"m:GetAttachment"`
}
