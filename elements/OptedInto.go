package elements

// The OptedInto element specifies a Boolean value that indicates whether the user opted in to the retention policy.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/optedinto
type OptedInto bool

const (
	// false
	OptedIntofalse OptedInto = false
	// true
	OptedIntotrue OptedInto = true
)
