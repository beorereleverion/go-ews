package elements

// The FindPeople element specifies a set of data used in a FindPeople request. The data includes zero or more of the following elements: a persona shape (optional), an indexed page item view, a restriction (optional), an aggregation restriction (optional), a sort order (optional), a parent folder Id, and a query string (optional).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople
import "encoding/xml"

type FindPeople struct {
	XMLName xml.Name

	// The AggregationRestriction element specifies a value that is applied to a set of Persona properties resulting from a FindPeople request and filters the result according to the specified restriction.
	AggregationRestriction *AggregationRestriction `xml:"AggregationRestriction"`
	// The IndexedPageItemView element describes how paged conversation or item information is returned for a FindItem operation or FindConversation operation request.
	IndexedPageItemView *IndexedPageItemView `xml:"IndexedPageItemView"`
	// The ParentFolderId element identifies the folder in which a new folder is created or the folder to search for the FindConversation operation.
	ParentFolderId *ParentFolderIdTargetFolderIdType `xml:"ParentFolderId"`
	// The PersonaShape element specifies the set of persona properties to be returned from a FindPeople request.
	PersonaShape *PersonaShape `xml:"PersonaShape"`
	// The QueryString element contains a mailbox query string based on Advanced Query Syntax (AQS).
	QueryString *QueryStringQueryStringType `xml:"QueryString"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"Restriction"`
	// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
	SortOrder *SortOrder `xml:"SortOrder"`
}

func (F *FindPeople) SetForMarshal() {
	F.XMLName.Local = "m:FindPeople"
}

func (F *FindPeople) GetSchema() *Schema {
	return &SchemaMessages
}
