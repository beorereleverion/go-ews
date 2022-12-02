package elements

// The IsVoicemail element indicates whether incoming messages must be voice mail messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isvoicemail
type IsVoicemail bool

const (
	// false
	IsVoicemailfalse IsVoicemail = false
	// true
	IsVoicemailtrue IsVoicemail = true
)
