package elements

// The MailboxSearchScopes element specifies a list of one or more mailboxes and associated search scopes for a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxsearchscopes
type MailboxSearchScopes struct {
	// The MailboxSearchScope element specifies a mailbox and a search scope for a discovery search.
	MailboxSearchScope *MailboxSearchScope `xml:"t:MailboxSearchScope"`
}
