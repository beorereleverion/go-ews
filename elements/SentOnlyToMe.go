package elements

// The SentOnlyToMe element indicates whether the owner of the mailbox has to be the only one in the ToRecipients property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sentonlytome
type SentOnlyToMe bool

const (
	// false
	SentOnlyToMefalse SentOnlyToMe = false
	// true
	SentOnlyToMetrue SentOnlyToMe = true
)
