package elements

// The Value element specifies an array of extended properties for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-extendedpropertytype
import "encoding/xml"

type ValueExtendedPropertyType struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
	// The Values element contains a collection of values for an extended property.
	Values *Values `xml:"Values"`
}

func (V *ValueExtendedPropertyType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueExtendedPropertyType) GetSchema() *Schema {
	return &SchemaTypes
}
