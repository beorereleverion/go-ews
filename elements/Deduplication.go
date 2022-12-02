package elements

// The Deduplication element indicates whether the search result should remove duplicate items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deduplication
type Deduplication bool

const (
	// false
	Deduplicationfalse Deduplication = false
	// true
	Deduplicationtrue Deduplication = true
)
