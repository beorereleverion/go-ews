package elements

// The MiddleNames element specifies an array of middle name values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/middlenames
import "encoding/xml"

type MiddleNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (M *MiddleNames) SetForMarshal() {
	M.XMLName.Local = "t:MiddleNames"
}

func (M *MiddleNames) GetSchema() *Schema {
	return &SchemaTypes
}
