package elements

// The SearchQueries element contains a list of mailboxes and associated queries for discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchqueries
type SearchQueries struct {
	// The MailboxQuery element specifies a query and the scope of a discovery search.
	MailboxQuery *MailboxQuery `xml:"t:MailboxQuery"`
}
