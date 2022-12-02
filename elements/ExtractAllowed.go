package elements

// The ExtractAllowed element specifies whether entity extraction is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extractallowed
type ExtractAllowed bool

const (
	// false
	ExtractAllowedfalse ExtractAllowed = false
	// true
	ExtractAllowedtrue ExtractAllowed = true
)
