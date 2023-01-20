package elements

// The Alias element contains the email alias of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/alias
import "encoding/xml"

type Alias struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *Alias) SetForMarshal() {
	A.XMLName.Local = "t:Alias"
}

func (A *Alias) GetSchema() *Schema {
	return &SchemaTypes
}
