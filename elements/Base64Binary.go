package elements

// The Base64Binary element contains a Base64-encoded value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/base64binary
import "encoding/xml"

type Base64Binary struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *Base64Binary) SetForMarshal() {
	B.XMLName.Local = "t:Base64Binary"
}

func (B *Base64Binary) GetSchema() *Schema {
	return &SchemaTypes
}
