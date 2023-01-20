package elements

// The PhoneticFirstName element contains the first name of a contact, spelled phonetically.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phoneticfirstname
import "encoding/xml"

type PhoneticFirstName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PhoneticFirstName) SetForMarshal() {
	P.XMLName.Local = "t:PhoneticFirstName"
}

func (P *PhoneticFirstName) GetSchema() *Schema {
	return &SchemaTypes
}
