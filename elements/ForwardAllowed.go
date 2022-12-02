package elements

// The ForwardAllowed element specifies whether forwarding emails is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardallowed
type ForwardAllowed bool

const (
	// false
	ForwardAllowedfalse ForwardAllowed = false
	// true
	ForwardAllowedtrue ForwardAllowed = true
)
