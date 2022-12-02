package elements

// The PolicyTipsEnabled element indicates whether policy tips are enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/policytipsenabled
type PolicyTipsEnabled bool

const (
	// false
	PolicyTipsEnabledfalse PolicyTipsEnabled = false
	// true
	PolicyTipsEnabledtrue PolicyTipsEnabled = true
)
