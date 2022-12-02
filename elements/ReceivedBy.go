package elements

// The ReceivedBy element identifies the delegate in a delegate access scenario.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedby
type ReceivedBy struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
