package elements

// The Initials element represents the initials of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/initials
import "encoding/xml"

type Initials struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *Initials) SetForMarshal() {
	I.XMLName.Local = "t:Initials"
}

func (I *Initials) GetSchema() *Schema {
	return &SchemaTypes
}
