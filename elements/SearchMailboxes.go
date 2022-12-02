package elements

// The SearchMailboxes element indicates the beginning of the SearchMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxes
type SearchMailboxes struct {
	// The Deduplication element indicates whether the search result should remove duplicate items.
	Deduplication *Deduplication `xml:"t:Deduplication"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"m:Language"`
	// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
	PageDirection *PageDirection `xml:"m:PageDirection"`
	// The PageItemReference element specifies the reference for a page item.
	PageItemReference *PageItemReference `xml:"m:PageItemReference"`
	// The PageSize element contains the number of items to be returned in a single page for a search result.
	PageSize *PageSize `xml:"m:PageSize"`
	// The PreviewItemResponseShape element contains the requested property set to be returned in a discovery search.
	PreviewItemResponseShape *PreviewItemResponseShape `xml:"m:PreviewItemResponseShape"`
	// The ResultType element contains the type of search to perform. The type of search can be statistics only or preview only.
	ResultType *ResultType `xml:"m:ResultType"`
	// The SearchQueries element contains a list of mailboxes and associated queries for discovery search.
	SearchQueries *SearchQueries `xml:"m:SearchQueries"`
	// The SortBy element contains an item property used for sorting the search result.
	SortBy *SortBy `xml:"m:SortBy"`
}
