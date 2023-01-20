package elements

// The GivenNames element specifies an array of given name values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/givennames
import "encoding/xml"

type GivenNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (G *GivenNames) SetForMarshal() {
	G.XMLName.Local = "t:GivenNames"
}

func (G *GivenNames) GetSchema() *Schema {
	return &SchemaTypes
}
