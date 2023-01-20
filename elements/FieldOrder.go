package elements

// The FieldOrder element represents a single field by which to sort results and indicates the direction for the sort.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fieldorder
import "encoding/xml"

type FieldOrder struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// Describes the sort order direction. The following are the possible values:  - Ascending  - Descending
	Order *string `xml:"Order,attr"`
}

func (F *FieldOrder) SetForMarshal() {
	F.XMLName.Local = "t:FieldOrder"
}

func (F *FieldOrder) GetSchema() *Schema {
	return &SchemaTypes
}
