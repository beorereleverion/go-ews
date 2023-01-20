package elements

// The HomeCity element specifies the city of the home address of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homecity
import "encoding/xml"

type HomeCity struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (H *HomeCity) SetForMarshal() {
	H.XMLName.Local = "t:HomeCity"
}

func (H *HomeCity) GetSchema() *Schema {
	return &SchemaTypes
}
