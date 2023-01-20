package elements

// The ExtendedProperties element contains the extended properties used for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-arrayofextendedpropertyattributedvaluetype
import "encoding/xml"

type ExtendedPropertiesArrayOfExtendedPropertyAttributedValueType struct {
	XMLName xml.Name

	// The ExtendedPropertyAttributedValue element specifies extended properties for a persona.
	ExtendedPropertyAttributedValue *ExtendedPropertyAttributedValue `xml:"ExtendedPropertyAttributedValue"`
}

func (E *ExtendedPropertiesArrayOfExtendedPropertyAttributedValueType) SetForMarshal() {
	E.XMLName.Local = "t:ExtendedProperties"
}

func (E *ExtendedPropertiesArrayOfExtendedPropertyAttributedValueType) GetSchema() *Schema {
	return &SchemaTypes
}
