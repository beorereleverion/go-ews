package elements

// The SearchMailboxesResult element contains the result of the SearchMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxesresult
type SearchMailboxesResult struct {
	// The FailedMailboxes element specifies an array of mailboxes that failed on search.
	FailedMailboxes *FailedMailboxes `xml:"t:FailedMailboxes"`
	// The ItemCount element specifies the total number of items in a search result.
	ItemCount *ItemCount `xml:"ItemCount"`
	// The Items element specifies a list of items available for preview as the results of a SearchMailboxes operation.
	Items *ItemsArrayOfSearchPreviewItemsType `xml:"Items"`
	// The KeywordStats element specifies a list of one or more KeywordStat elements.
	KeywordStats *KeywordStats `xml:"KeywordStats"`
	// The MailboxStats element specifies a list of one or more MailboxStat elements.
	MailboxStats *MailboxStats `xml:"t:MailboxStats"`
	// The PageItemCount element specifies the number of pages returned in a search result pagination.
	PageItemCount *PageItemCount `xml:"t:PageItemCount"`
	// The PageItemSize element specifies the number of items to return in a search result pagination.
	PageItemSize *PageItemSize `xml:"t:PageItemSize"`
	// The Refiners element specifies a list of one or more Refiner elements.
	Refiners *Refiners `xml:"t:Refiners"`
	// The ResultType element contains the type of search to perform. The type of search can be statistics only or preview only.
	ResultType *ResultType `xml:"m:ResultType"`
	// The SearchQueries element contains a list of mailboxes and associated queries for discovery search.
	SearchQueries *SearchQueries `xml:"m:SearchQueries"`
	// The Size element specifies the total size of one or more mailbox items.
	Size *Sizelong `xml:"t:Size"`
}
