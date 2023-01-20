package elements

// The GivenName element contains a contact's given name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/givenname
import "encoding/xml"

type GivenName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *GivenName) SetForMarshal() {
	G.XMLName.Local = "t:GivenName"
}

func (G *GivenName) GetSchema() *Schema {
	return &SchemaTypes
}
