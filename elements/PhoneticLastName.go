package elements

// The PhoneticLastName element contains the last name of a contact, spelled phonetically.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phoneticlastname
import "encoding/xml"

type PhoneticLastName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PhoneticLastName) SetForMarshal() {
	P.XMLName.Local = "t:PhoneticLastName"
}

func (P *PhoneticLastName) GetSchema() *Schema {
	return &SchemaTypes
}
