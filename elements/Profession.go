package elements

// The Profession element represents the profession of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/profession
import "encoding/xml"

type Profession struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *Profession) SetForMarshal() {
	P.XMLName.Local = "t:Profession"
}

func (P *Profession) GetSchema() *Schema {
	return &SchemaTypes
}
