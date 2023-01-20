package elements

// The ExtendedProperty element identifies extended MAPI properties on folders and items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperty
import "encoding/xml"

type ExtendedProperty struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
	// The Values element contains a collection of values for an extended property.
	Values *Values `xml:"Values"`
}

func (E *ExtendedProperty) SetForMarshal() {
	E.XMLName.Local = "t:ExtendedProperty"
}

func (E *ExtendedProperty) GetSchema() *Schema {
	return &SchemaTypes
}
