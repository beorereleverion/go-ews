package elements

// The Value element contains the value of an extended property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value
import "encoding/xml"

type Value struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *Value) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *Value) GetSchema() *Schema {
	return &SchemaTypes
}
