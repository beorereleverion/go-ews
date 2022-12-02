package elements

// The MailboxQuery element specifies a query and the scope of a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxquery
type MailboxQuery struct {
	// The MailboxSearchScopes element specifies a list of one or more mailboxes and associated search scopes for a discovery search.
	MailboxSearchScopes *MailboxSearchScopes `xml:"t:MailboxSearchScopes"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"m:Query"`
}
