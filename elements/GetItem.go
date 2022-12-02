package elements

// The GetItem element defines a request to get an item from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitem
type GetItem struct {
	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"m:ItemIds"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"m:ItemShape"`
}
