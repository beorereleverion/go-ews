package elements

// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemshape
type ItemShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"t:BaseShape"`
	// The BodyType element identifies how the body text is formatted in the response.
	BodyType *BodyType `xml:"t:BodyType"`
	// The ConvertHtmlCodePageToUTF8 element indicates whether the item HTML body is converted to UTF8.
	ConvertHtmlCodePageToUTF8 *ConvertHtmlCodePageToUTF8 `xml:"t:ConvertHtmlCodePageToUTF8"`
	// The FilterHtmlContent element specifies whether potentially unsafe HTML content is filtered from an item or attachment.
	FilterHtmlContent *FilterHtmlContent `xml:"t:FilterHtmlContent"`
	// The IncludeMimeContent element specifies whether the Multipurpose Internet Mail Extensions (MIME) content of an item or attachment is returned in the response.
	IncludeMimeContent *IncludeMimeContent `xml:"t:IncludeMimeContent"`
}
