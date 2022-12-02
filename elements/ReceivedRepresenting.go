package elements

// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedrepresenting
type ReceivedRepresenting struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
