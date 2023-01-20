package elements

// The PhoneticFullName element contains the full name of a contact, including the first and last name, spelled phonetically.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phoneticfullname
import "encoding/xml"

type PhoneticFullName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PhoneticFullName) SetForMarshal() {
	P.XMLName.Local = "t:PhoneticFullName"
}

func (P *PhoneticFullName) GetSchema() *Schema {
	return &SchemaTypes
}
