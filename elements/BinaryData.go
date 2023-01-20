package elements

// The BinaryData element contains binary data property content.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/binarydata
import "encoding/xml"

type BinaryData struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BinaryData) SetForMarshal() {
	B.XMLName.Local = "t:BinaryData"
}

func (B *BinaryData) GetSchema() *Schema {
	return &SchemaTypes
}
