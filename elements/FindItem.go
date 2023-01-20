package elements

// The FindItem element defines a request to find items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditem
import "encoding/xml"

type FindItem struct {
	XMLName xml.Name

	// The CalendarView element defines a FindItem operation as returning calendar items in a set as they appear in a calendar.
	CalendarView *CalendarView `xml:"CalendarView"`
	// The ContactsView element defines a search for contact items based on alphabetical display names.
	ContactsView *ContactsView `xml:"ContactsView"`
	// The DistinguishedGroupBy element provides standard groupings for FindItem queries.
	DistinguishedGroupBy *DistinguishedGroupBy `xml:"DistinguishedGroupBy"`
	// The FractionalPageItemView element describes where the paged view starts and the maximum number of items returned in a FindItem request.
	FractionalPageItemView *FractionalPageItemView `xml:"FractionalPageItemView"`
	// The GroupBy element specifies an arbitrary grouping for FindItem queries.
	GroupBy *GroupBy `xml:"GroupBy"`
	// The IndexedPageItemView element describes how paged conversation or item information is returned for a FindItem operation or FindConversation operation request.
	IndexedPageItemView *IndexedPageItemView `xml:"IndexedPageItemView"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"ItemShape"`
	// The ParentFolderIds element identifies folders for the FindItem and FindFolder operations to search.
	ParentFolderIds *ParentFolderIds `xml:"ParentFolderIds"`
	// The QueryString element contains a mailbox query string based on Advanced Query Syntax (AQS).
	QueryString *QueryStringQueryStringType `xml:"QueryString"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"Restriction"`
	// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
	SortOrder *SortOrder `xml:"SortOrder"`
	// Defines whether the search finds items in folders or the folders' dumpsters. This attribute is required.
	Traversal *string `xml:"Traversal,attr"`
}

const (
	// Returns only the identities of items in the folder.
	FindItemShallow = `Shallow`
	// Returns only the identities of items that are in a folder's dumpster. Note that a soft-deleted traversal combined with a search restriction will result in zero items returned even if there are items that match the search criteria.
	FindItemSoftDeleted = `SoftDeleted`
	// Returns only the identities of associated items in the folder.
	FindItemAssociated = `Associated`
)

func (F *FindItem) SetForMarshal() {
	F.XMLName.Local = "m:FindItem"
}

func (F *FindItem) GetSchema() *Schema {
	return &SchemaMessages
}
