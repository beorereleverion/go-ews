package elements

// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/torecipients
type ToRecipients struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
