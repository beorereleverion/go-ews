package elements

// The HasLocationChanged element specifies whether the location property of a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haslocationchanged
type HasLocationChanged bool

const (
	// false
	HasLocationChangedfalse HasLocationChanged = false
	// true
	HasLocationChangedtrue HasLocationChanged = true
)
