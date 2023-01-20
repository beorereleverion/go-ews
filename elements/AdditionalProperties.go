package elements

// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additionalproperties
import "encoding/xml"

type AdditionalProperties struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (A *AdditionalProperties) SetForMarshal() {
	A.XMLName.Local = "t:AdditionalProperties"
}

func (A *AdditionalProperties) GetSchema() *Schema {
	return &SchemaTypes
}
