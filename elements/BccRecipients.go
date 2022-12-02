package elements

// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bccrecipients
type BccRecipients struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
