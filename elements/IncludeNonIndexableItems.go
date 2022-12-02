package elements

// The IncludeNonIndexableItems element contains a Boolean value to indicate whether to include items that cannot be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includenonindexableitems
type IncludeNonIndexableItems bool

const (
	// false
	IncludeNonIndexableItemsfalse IncludeNonIndexableItems = false
	// true
	IncludeNonIndexableItemstrue IncludeNonIndexableItems = true
)
