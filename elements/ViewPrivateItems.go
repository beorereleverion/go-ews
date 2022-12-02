package elements

// The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/viewprivateitems
type ViewPrivateItems bool

const (
	// false
	ViewPrivateItemsfalse ViewPrivateItems = false
	// true
	ViewPrivateItemstrue ViewPrivateItems = true
)
