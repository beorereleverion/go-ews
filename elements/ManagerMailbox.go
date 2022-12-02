package elements

// The ManagerMailbox element contains SMTP information that identifies the mailbox of the contact's manager.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managermailbox
type ManagerMailbox struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
