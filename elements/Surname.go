package elements

// The Surname element represents the surname of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/surname
import "encoding/xml"

type Surname struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Surname) SetForMarshal() {
	S.XMLName.Local = "t:Surname"
}

func (S *Surname) GetSchema() *Schema {
	return &SchemaTypes
}
