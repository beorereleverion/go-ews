package elements

// The ExtendedPropertyAttributedValue element specifies extended properties for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedpropertyattributedvalue
import "encoding/xml"

type ExtendedPropertyAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies an array of extended properties for a persona.
	Value *ValueExtendedPropertyType `xml:"Value"`
}

func (E *ExtendedPropertyAttributedValue) SetForMarshal() {
	E.XMLName.Local = "t:ExtendedPropertyAttributedValue"
}

func (E *ExtendedPropertyAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
