package elements

// The Nickname element represents the nickname of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nickname
import "encoding/xml"

type Nickname struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *Nickname) SetForMarshal() {
	N.XMLName.Local = "t:Nickname"
}

func (N *Nickname) GetSchema() *Schema {
	return &SchemaTypes
}
