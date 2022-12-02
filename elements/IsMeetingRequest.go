package elements

// The IsMeetngRequest element indicates whether incoming messages must be a meeting request in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeetingrequest
type IsMeetingRequest bool

const (
	// false
	IsMeetingRequestfalse IsMeetingRequest = false
	// true
	IsMeetingRequesttrue IsMeetingRequest = true
)
