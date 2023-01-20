package elements

// The IndexedPageFolderView element describes how paged item information is returned in a FindFolder response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedpagefolderview
import "encoding/xml"

type IndexedPageFolderView struct {
	XMLName xml.Name

	// Describes whether the page of folders will start from the start or the end of the set of folders that are found with the search criteria. Seeking from the end always searches backward. This attribute is required.
	BasePoint *string `xml:"BasePoint,attr"`
	// Describes the maximum number of folders to return in the response. This attribute is optional.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Describes the offset from the BasePoint. Offset must be greater than or equal to zero. If BasePoint is equal to Beginning, the offset is positive. If BasePoint is equal to End, the offset is handled as if it were negative.   This identifies which folder will be the first folder delivered in the response. This attribute is required.
	Offset *string `xml:"Offset,attr"`
}

func (I *IndexedPageFolderView) SetForMarshal() {
	I.XMLName.Local = "m:IndexedPageFolderView"
}

func (I *IndexedPageFolderView) GetSchema() *Schema {
	return &SchemaMessages
}
