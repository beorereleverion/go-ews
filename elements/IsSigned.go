package elements

// The IsSigned element indicates whether incoming messages must be signed in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/issigned
type IsSigned bool

const (
	// false
	IsSignedfalse IsSigned = false
	// true
	IsSignedtrue IsSigned = true
)
