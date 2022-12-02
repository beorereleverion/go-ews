package elements

// The IsNDR element indicates whether incoming messages must be non-delivery reports (NDRs) in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isndr
type IsNDR bool

const (
	// false
	IsNDRfalse IsNDR = false
	// true
	IsNDRtrue IsNDR = true
)
