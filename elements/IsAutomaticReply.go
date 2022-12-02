package elements

// The IsAutomaticReply element indicates whether incoming messages must be automatic replies in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isautomaticreply
type IsAutomaticReply bool

const (
	// false
	IsAutomaticReplyfalse IsAutomaticReply = false
	// true
	IsAutomaticReplytrue IsAutomaticReply = true
)
