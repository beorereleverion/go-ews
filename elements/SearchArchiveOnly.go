package elements

// The SearchArchiveOnly element indicates whether only the archive mailbox is searched for non-indexable items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searcharchiveonly
type SearchArchiveOnly bool

const (
	// false
	SearchArchiveOnlyfalse SearchArchiveOnly = false
	// true
	SearchArchiveOnlytrue SearchArchiveOnly = true
)
