package elements

// The AttachmentShape element identifies additional properties to return in a response to a GetAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentshape
type AttachmentShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
	// The BodyType element identifies how the body text is formatted in the response.
	BodyType *BodyType `xml:"t:BodyType"`
	// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
	FilterHtmlContent *FilterHtmlContent `xml:"t:FilterHtmlContent"`
	// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
	IncludeMimeContent *IncludeMimeContent `xml:"t:IncludeMimeContent"`
}
