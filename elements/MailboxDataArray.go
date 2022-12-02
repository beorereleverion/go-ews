package elements

// The MailboxDataArray element contains a list of mailboxes to query for availability information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxdataarray
type MailboxDataArray struct {
	// The MailboxData element represents an individual mailbox user and options for the type of data to be returned about the mailbox user.
	MailboxData *MailboxData `xml:"t:MailboxData"`
}
