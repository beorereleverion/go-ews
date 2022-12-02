package elements

// The GetNonIndexableItemDetails element specifies a request to retrieve nonindexable item details.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemdetails
type GetNonIndexableItemDetails struct {
	// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
	Mailboxes *MailboxesNonEmptyArrayOfLegacyDNsType `xml:"m:Mailboxes"`
	// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
	PageDirection *PageDirection `xml:"m:PageDirection"`
	// The PageItemReference element specifies the reference for a page item.
	PageItemReference *PageItemReference `xml:"m:PageItemReference"`
	// The PageSize element contains the number of items to be returned in a single page for a search result.
	PageSize *PageSize `xml:"m:PageSize"`
}
