package elements

// The From element represents the address from which the message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/from
type From struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
