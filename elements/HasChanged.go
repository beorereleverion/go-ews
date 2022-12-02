package elements

// The HasChanged element indicates whether a user's photo has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haschanged
type HasChanged string

const (
	// false
	HasChangedfalse HasChanged = `false`
	// true
	HasChangedtrue HasChanged = `true`
)
