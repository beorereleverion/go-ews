package elements

// The BusinessHomePages element specifies an array of business home pages and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businesshomepages
import "encoding/xml"

type BusinessHomePages struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (B *BusinessHomePages) SetForMarshal() {
	B.XMLName.Local = "t:BusinessHomePages"
}

func (B *BusinessHomePages) GetSchema() *Schema {
	return &SchemaTypes
}
