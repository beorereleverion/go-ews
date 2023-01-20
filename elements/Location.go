package elements

// The Location element represents the location of a meeting, appointment, or persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/location
import "encoding/xml"

type Location struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *Location) SetForMarshal() {
	L.XMLName.Local = "t:Location"
}

func (L *Location) GetSchema() *Schema {
	return &SchemaTypes
}
