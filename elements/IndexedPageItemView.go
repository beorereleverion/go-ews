package elements

// The IndexedPageItemView element describes how paged conversation or item information is returned for a FindItem operation or FindConversation operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedpageitemview
import "encoding/xml"

type IndexedPageItemView struct {
	XMLName xml.Name

	// Describes whether the page of items or conversations will start from the beginning or the end of the set of items or conversations that are found by using the search criteria. Seeking from the end always searches backward. This attribute is required.
	BasePoint *string `xml:"BasePoint,attr"`
	// Describes the maximum number of items or conversations to return in the response. This attribute is optional.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Describes the offset from the BasePoint. If BasePoint is equal to Beginning, the offset is positive. If BasePoint is equal to End, the offset is handled as if it were negative. This identifies which item or conversation will be the first to be delivered in the response. This attribute is required.
	Offset *string `xml:"Offset,attr"`
}

func (I *IndexedPageItemView) SetForMarshal() {
	I.XMLName.Local = "m:IndexedPageItemView"
}

func (I *IndexedPageItemView) GetSchema() *Schema {
	return &SchemaMessages
}
