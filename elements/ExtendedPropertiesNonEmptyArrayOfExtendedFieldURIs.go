package elements

// The ExtendedProperties element contains the extended properties used for the Unified Contact Store operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperties-nonemptyarrayofextendedfielduris
import "encoding/xml"

type ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs struct {
	XMLName xml.Name

	// The ExtendedProperty element specifies an extended property for the Unified Contact Store.
	ExtendedProperty *ExtendedPropertyPathToExtendedFieldType `xml:"ExtendedProperty"`
}

func (E *ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs) SetForMarshal() {
	E.XMLName.Local = "m:ExtendedProperties"
}

func (E *ExtendedPropertiesNonEmptyArrayOfExtendedFieldURIs) GetSchema() *Schema {
	return &SchemaMessages
}
