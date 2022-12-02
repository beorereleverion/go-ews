package elements

// The DiscoverySearchConfiguration element specifies the configuration for eDiscovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/discoverysearchconfiguration
type DiscoverySearchConfiguration struct {
	// The SearchId element specifies the identifier of a discovery search.
	SearchId *SearchId `xml:"t:SearchId"`
	// The SearchQuery element specifies the discovery search query.
	SearchQuery *SearchQuery `xml:"t:SearchQuery"`
	// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
	SearchableMailboxes *SearchableMailboxes `xml:"m:SearchableMailboxes"`
}
