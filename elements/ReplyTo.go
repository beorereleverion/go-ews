package elements

// The ReplyTo element identifies an array of addresses to which replies should be sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyto
type ReplyTo struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
