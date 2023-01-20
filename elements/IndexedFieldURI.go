package elements

// The IndexedFieldURI element identifies individual members of a dictionary.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedfielduri
import "encoding/xml"

type IndexedFieldURI struct {
	XMLName xml.Name

	// Identifies the member of the dictionary to return. This attribute is required.
	FieldIndex *string `xml:"FieldIndex,attr"`
	// Identifies the dictionary that contains the member to return. This attribute is required.
	FieldURI *string `xml:"FieldURI,attr"`
}

func (I *IndexedFieldURI) SetForMarshal() {
	I.XMLName.Local = "t:IndexedFieldURI"
}

func (I *IndexedFieldURI) GetSchema() *Schema {
	return &SchemaTypes
}
