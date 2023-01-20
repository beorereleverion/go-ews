package elements

// The OfficeLocations element specifies an array of office locations and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/officelocations
import "encoding/xml"

type OfficeLocations struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (O *OfficeLocations) SetForMarshal() {
	O.XMLName.Local = "t:OfficeLocations"
}

func (O *OfficeLocations) GetSchema() *Schema {
	return &SchemaTypes
}
