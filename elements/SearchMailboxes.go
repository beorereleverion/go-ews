package elements

// The SearchMailboxes element indicates the beginning of the SearchMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxes
import "encoding/xml"

type SearchMailboxes struct {
	XMLName xml.Name

	// The Deduplication element indicates whether the search result should remove duplicate items.
	Deduplication *Deduplication `xml:"Deduplication"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"Language"`
	// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
	PageDirection *PageDirection `xml:"PageDirection"`
	// The PageItemReference element specifies the reference for a page item.
	PageItemReference *PageItemReference `xml:"PageItemReference"`
	// The PageSize element contains the number of items to be returned in a single page for a search result.
	PageSize *PageSize `xml:"PageSize"`
	// The PreviewItemResponseShape element contains the requested property set to be returned in a discovery search.
	PreviewItemResponseShape *PreviewItemResponseShape `xml:"PreviewItemResponseShape"`
	// The ResultType element contains the type of search to perform. The type of search can be statistics only or preview only.
	ResultType *ResultType `xml:"ResultType"`
	// The SearchQueries element contains a list of mailboxes and associated queries for discovery search.
	SearchQueries *SearchQueries `xml:"SearchQueries"`
	// The SortBy element contains an item property used for sorting the search result.
	SortBy *SortBy `xml:"SortBy"`
}

func (S *SearchMailboxes) SetForMarshal() {
	S.XMLName.Local = "m:SearchMailboxes"
}

func (S *SearchMailboxes) GetSchema() *Schema {
	return &SchemaMessages
}
