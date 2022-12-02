package elements

// The Mailboxes element contains an array of mailboxes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-arrayofusermailboxestype
type MailboxesArrayOfUserMailboxesType struct {
	// The UserMailbox element identifies a user mailbox.
	UserMailbox *UserMailbox `xml:"t:UserMailbox"`
}
