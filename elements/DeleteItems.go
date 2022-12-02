package elements

// The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitems
type DeleteItems string

const (
	// All
	DeleteItemsAll DeleteItems = `All`
	// None
	DeleteItemsNone DeleteItems = `None`
	// Owned
	DeleteItemsOwned DeleteItems = `Owned`
)
