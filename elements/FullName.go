package elements

// The FullName element represents the full name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fullname
import "encoding/xml"

type FullName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FullName) SetForMarshal() {
	F.XMLName.Local = "t:FullName"
}

func (F *FullName) GetSchema() *Schema {
	return &SchemaTypes
}
