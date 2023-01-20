package elements

// The ImAddresses3 element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddresses3
import "encoding/xml"

type ImAddresses3 struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (I *ImAddresses3) SetForMarshal() {
	I.XMLName.Local = "t:ImAddresses3"
}

func (I *ImAddresses3) GetSchema() *Schema {
	return &SchemaTypes
}
