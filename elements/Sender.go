package elements

// The Sender element identifies the sender of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender
type Sender struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
