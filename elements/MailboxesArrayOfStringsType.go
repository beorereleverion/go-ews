package elements

// The Mailboxes element contains a list of mailboxes affected by the hold.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-arrayofstringstype
type MailboxesArrayOfStringsType struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
