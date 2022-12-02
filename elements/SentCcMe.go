package elements

// The SentCcMe element indicates whether the owner of the mailbox has to be in the CcRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sentccme
type SentCcMe bool

const (
	// false
	SentCcMefalse SentCcMe = false
	// true
	SentCcMetrue SentCcMe = true
)
