package elements

// The Altitude element specifies the altitude of a postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/altitude
import "encoding/xml"

type Altitude struct {
	XMLName xml.Name
	TEXT    float64 `xml:",chardata"`
}

func (A *Altitude) SetForMarshal() {
	A.XMLName.Local = "t:Altitude"
}

func (A *Altitude) GetSchema() *Schema {
	return &SchemaTypes
}
