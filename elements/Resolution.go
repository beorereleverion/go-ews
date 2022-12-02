package elements

// The Resolution element contains a single resolved entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolution
type Resolution struct {
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"t:Contact"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
