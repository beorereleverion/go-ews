package elements

// The Intent element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/intent
import "encoding/xml"

type Intent struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (I *Intent) SetForMarshal() {
	I.XMLName.Local = "t:Intent"
}

func (I *Intent) GetSchema() *Schema {
	return &SchemaTypes
}
