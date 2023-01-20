package elements

// The ExtendedProperties element specifies an array of additional properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-nonemptyarrayofextendedpropertytype
import "encoding/xml"

type ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType struct {
	XMLName xml.Name

	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"ExtendedProperty"`
}

func (E *ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType) SetForMarshal() {
	E.XMLName.Local = "t:ExtendedProperties"
}

func (E *ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType) GetSchema() *Schema {
	return &SchemaTypes
}
