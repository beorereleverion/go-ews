package elements

// The PrintAllowed element specifies whether printing is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/printallowed
type PrintAllowed bool

const (
	// false
	PrintAllowedfalse PrintAllowed = false
	// true
	PrintAllowedtrue PrintAllowed = true
)
