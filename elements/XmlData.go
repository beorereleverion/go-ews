package elements

// The XmlData element contains XML data property content for a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/xmldata
import "encoding/xml"

type XmlData struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (X *XmlData) SetForMarshal() {
	X.XMLName.Local = "t:XmlData"
}

func (X *XmlData) GetSchema() *Schema {
	return &SchemaTypes
}
