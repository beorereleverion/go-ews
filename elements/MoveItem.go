package elements

// The MoveItem element defines a request to move an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moveitem
type MoveItem struct {
	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"m:ItemIds"`
	// The ReturnNewItemIds element indicates whether the item identifiers of new items are returned in the response.
	ReturnNewItemIds *ReturnNewItemIds `xml:"m:ReturnNewItemIds"`
	// The ToFolderId element represents the destination folder for a copied or moved item or folder.
	ToFolderId *ToFolderId `xml:"m:ToFolderId"`
}
