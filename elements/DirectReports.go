package elements

// The DirectReports element contains SMTP information that identifies the direct reports of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/directreports
type DirectReports struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
