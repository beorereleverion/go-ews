package elements

// The SearchScope element specifies the scope of a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchscope
type SearchScope string

const (
	// All
	SearchScopeAll SearchScope = `All`
	// ArchiveOnly
	SearchScopeArchiveOnly SearchScope = `ArchiveOnly`
	// PrimaryOnly
	SearchScopePrimaryOnly SearchScope = `PrimaryOnly`
)
