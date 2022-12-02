package elements

// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ccrecipients
type CcRecipients struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
