package elements

// The Preview element specifies the first 256 characters of an item for display.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/preview-ex15websvcsotherref
import "encoding/xml"

type Preview struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *Preview) SetForMarshal() {
	P.XMLName.Local = "t:Preview"
}

func (P *Preview) GetSchema() *Schema {
	return &SchemaTypes
}
