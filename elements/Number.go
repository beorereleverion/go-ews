package elements

// The Number element specifies a phone number.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/number
import "encoding/xml"

type Number struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *Number) SetForMarshal() {
	N.XMLName.Local = "t:Number"
}

func (N *Number) GetSchema() *Schema {
	return &SchemaTypes
}
