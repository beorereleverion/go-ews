package elements

// The UmEnabled element indicates whether Unified Messaging is enabled for an account.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/umenabled
type UmEnabled bool

const (
	// false
	UmEnabledfalse UmEnabled = false
	// true
	UmEnabledtrue UmEnabled = true
)
