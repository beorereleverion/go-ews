package elements

// The MailboxSearchScope element specifies a mailbox and a search scope for a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxsearchscope
type MailboxSearchScope struct {
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
	// The SearchScope element specifies the scope of a search.
	SearchScope *SearchScope `xml:"t:SearchScope"`
}
