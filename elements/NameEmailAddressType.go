package elements

// The Name element represents the name of a mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-emailaddresstype
import "encoding/xml"

type NameEmailAddressType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NameEmailAddressType) SetForMarshal() {
	N.XMLName.Local = "t:Name"
}

func (N *NameEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
