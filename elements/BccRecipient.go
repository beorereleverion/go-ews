package elements

// The BccRecipient element represents a recipient to receive a blind carbon copy (Bcc) of an e-mail message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bccrecipient
type BccRecipient bool

const (
	// false
	BccRecipientfalse BccRecipient = false
	// true
	BccRecipienttrue BccRecipient = true
)
