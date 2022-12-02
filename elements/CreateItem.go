package elements

// The CreateItem element defines a request to create an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem
type CreateItem struct {
	// The Items element contains a set of items to create.
	Items *ItemsNonEmptyArrayOfAllItemsType `xml:"m:Items"`
	// The SavedItemFolderId element identifies the target folder for operations that update, send, and create items in a mailbox.
	SavedItemFolderId *SavedItemFolderId `xml:"m:SavedItemFolderId"`
	// Describes how the item will be handled after it is created. The attribute is required for e-mail messages. This attribute is only applicable to e-mail messages.
	MessageDisposition *string `xml:"MessageDisposition,attr"`
	// Describes how meeting requests are handled after they are created. This attribute is required for calendar items.
	SendMeetingInvitations *string `xml:"SendMeetingInvitations,attr"`
}
