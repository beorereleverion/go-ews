package elements

// The City element represents the city name that is associated with a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/city
import "encoding/xml"

type City struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *City) SetForMarshal() {
	C.XMLName.Local = "t:City"
}

func (C *City) GetSchema() *Schema {
	return &SchemaTypes
}
