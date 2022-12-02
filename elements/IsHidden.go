package elements

// The IsHidden element contains a Boolean value that indicates whether the underlying contact should be hidden or displayed as part of the persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ishidden
type IsHidden bool

const (
	// false
	IsHiddenfalse IsHidden = false
	// true
	IsHiddentrue IsHidden = true
)
