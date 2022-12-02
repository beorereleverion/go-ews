package elements

// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchablemailboxes
type SearchableMailboxes struct {
	// The SearchableMailbox element specifies a mailbox returned from a GetSearchableMailboxes request.
	SearchableMailbox *SearchableMailbox `xml:"t:SearchableMailbox"`
}
