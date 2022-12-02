package elements

// The Organizer element represents the organizer of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/organizer
type Organizer struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
