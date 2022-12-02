package elements

// The IsPartiallyIndexed element indicates whether the item is partially indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispartiallyindexed
type IsPartiallyIndexed bool

const (
	// false
	IsPartiallyIndexedfalse IsPartiallyIndexed = false
	// true
	IsPartiallyIndexedtrue IsPartiallyIndexed = true
)
