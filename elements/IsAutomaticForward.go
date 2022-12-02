package elements

// The IsAutomaticForward element indicates whether incoming messages must be automatic forwards in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isautomaticforward
type IsAutomaticForward bool

const (
	// false
	IsAutomaticForwardfalse IsAutomaticForward = false
	// true
	IsAutomaticForwardtrue IsAutomaticForward = true
)
