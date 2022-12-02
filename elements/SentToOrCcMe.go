package elements

// The SentToOrCcMe element indicates whether the owner of the mailbox has to be in either a ToRecipients or CcRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttoorccme
type SentToOrCcMe bool

const (
	// false
	SentToOrCcMefalse SentToOrCcMe = false
	// true
	SentToOrCcMetrue SentToOrCcMe = true
)
