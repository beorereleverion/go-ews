package elements

// The Address element specifies the address of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/address-contacttype
import "encoding/xml"

type AddressContactType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *AddressContactType) SetForMarshal() {
	A.XMLName.Local = "t:Address"
}

func (A *AddressContactType) GetSchema() *Schema {
	return &SchemaTypes
}
