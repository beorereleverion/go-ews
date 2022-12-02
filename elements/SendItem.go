package elements

// The SendItem element is the root element in a request to send an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senditem
type SendItem struct {
	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"m:ItemIds"`
	// The SavedItemFolderId element identifies the target folder for operations that update, send, and create items in a mailbox.
	SavedItemFolderId *SavedItemFolderId `xml:"m:SavedItemFolderId"`
	// Identifies whether a copy of the sent item is saved. The save action depends on the value of SaveItemToFolder and whether a SavedItemFolderId element is present in the request. This element is required.
	SaveItemToFolder *string `xml:"SaveItemToFolder,attr"`
}
