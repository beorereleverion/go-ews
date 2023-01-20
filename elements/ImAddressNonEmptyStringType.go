package elements

// The ImAddress element contains the instant messaging address of a new contact that will be added to an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddress-nonemptystringtype
import "encoding/xml"

type ImAddressNonEmptyStringType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ImAddressNonEmptyStringType) SetForMarshal() {
	I.XMLName.Local = "m:ImAddress"
}

func (I *ImAddressNonEmptyStringType) GetSchema() *Schema {
	return &SchemaMessages
}
