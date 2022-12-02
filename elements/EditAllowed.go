package elements

// The EditAllowed element specifies whether Information Rights Management can be edited.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/editallowed
type EditAllowed bool

const (
	// false
	EditAllowedfalse EditAllowed = false
	// true
	EditAllowedtrue EditAllowed = true
)
