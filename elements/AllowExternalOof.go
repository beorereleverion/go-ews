package elements

// The AllowExternalOof element contains a value that identifies to whom external Out of Office (OOF) messages are sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/allowexternaloof
type AllowExternalOof string

const (
	// E-mail senders outside the mailbox user's organization who send messages to the user will receive an external OOF message response.
	AllowExternalOofAll AllowExternalOof = `All`
	// E-mail senders outside the mailbox user's organization who send messages to the user will only receive an external OOF message response if the sender is in the user's Exchange store contact list.
	AllowExternalOofKnown AllowExternalOof = `Known`
	// E-mail senders outside the mailbox user's organization who send messages to the user will not receive an external OOF message response.
	AllowExternalOofNone AllowExternalOof = `None`
)
