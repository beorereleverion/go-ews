package elements

// The IncludeUnsearchableItems element specifies whether to include items that cannot be searched.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includeunsearchableitems
type IncludeUnsearchableItems bool

const (
	// false
	IncludeUnsearchableItemsfalse IncludeUnsearchableItems = false
	// true
	IncludeUnsearchableItemstrue IncludeUnsearchableItems = true
)
