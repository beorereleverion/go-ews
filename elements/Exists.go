package elements

// The Exists element represents a search expression that returns true if the supplied property exists on an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exists
import "encoding/xml"

type Exists struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (E *Exists) SetForMarshal() {
	E.XMLName.Local = "t:Exists"
}

func (E *Exists) GetSchema() *Schema {
	return &SchemaTypes
}
