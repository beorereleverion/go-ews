package elements

// The LastName element represents the last name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastname
import "encoding/xml"

type LastName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LastName) SetForMarshal() {
	L.XMLName.Local = "t:LastName"
}

func (L *LastName) GetSchema() *Schema {
	return &SchemaTypes
}
