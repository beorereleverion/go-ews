package elements

// The GetSearchableMailboxes element contains a request to get a list of mailboxes that the client has permission to perform an eDiscovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsearchablemailboxes
type GetSearchableMailboxes struct {
	// The ExpandGroupMembership element indicates whether to expand the membership of the group returned from a GetSearchableMailboxes request.
	ExpandGroupMembership *ExpandGroupMembership `xml:"m:ExpandGroupMembership"`
	// The SearchFilter element contains the query string to filter the mailboxes to be returned from a GetSearchableMailboxes request.
	SearchFilter *SearchFilter `xml:"m:SearchFilter"`
}
