package elements

// The NotSentToMe element indicates whether the owner of the mailbox must not be in the ToRecipients property of the incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notsenttome
type NotSentToMe bool

const (
	// false
	NotSentToMefalse NotSentToMe = false
	// true
	NotSentToMetrue NotSentToMe = true
)
