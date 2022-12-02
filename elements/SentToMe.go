package elements

// The SentToMe element indicates whether the owner of the mailbox has to be in the ToRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttome
type SentToMe bool

const (
	// false
	SentToMefalse SentToMe = false
	// true
	SentToMetrue SentToMe = true
)
