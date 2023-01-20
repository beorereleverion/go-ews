package elements

// The Generation element represents a generational abbreviation that follows the full name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/generation
import "encoding/xml"

type Generation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *Generation) SetForMarshal() {
	G.XMLName.Local = "t:Generation"
}

func (G *Generation) GetSchema() *Schema {
	return &SchemaTypes
}
