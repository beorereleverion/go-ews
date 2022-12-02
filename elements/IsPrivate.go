package elements

// The IsPrivate element indicates whether the calendar item is private.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isprivate
type IsPrivate bool

const (
	// false
	IsPrivatefalse IsPrivate = false
	// true
	IsPrivatetrue IsPrivate = true
)
