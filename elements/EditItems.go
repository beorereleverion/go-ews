package elements

// The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/edititems
type EditItems string

const (
	// All
	EditItemsAll EditItems = `All`
	// None
	EditItemsNone EditItems = `None`
	// Owned
	EditItemsOwned EditItems = `Owned`
)
