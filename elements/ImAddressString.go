package elements

// The ImAddress element contains the primary instant messaging address of a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddress-string
import "encoding/xml"

type ImAddressString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ImAddressString) SetForMarshal() {
	I.XMLName.Local = "t:ImAddress"
}

func (I *ImAddressString) GetSchema() *Schema {
	return &SchemaTypes
}
