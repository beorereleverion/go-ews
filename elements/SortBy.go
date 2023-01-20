package elements

// The SortBy element contains an item property used for sorting the search result.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortby
import "encoding/xml"

type SortBy struct {
	XMLName xml.Name

	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// The text value of the Order attribute is the sort order. A text value of Ascending indicates that the results are in ascending order. A text value of Descending indicates that the results are in descending order.
	Order *string `xml:"Order,attr"`
}

func (S *SortBy) SetForMarshal() {
	S.XMLName.Local = "m:SortBy"
}

func (S *SortBy) GetSchema() *Schema {
	return &SchemaMessages
}
