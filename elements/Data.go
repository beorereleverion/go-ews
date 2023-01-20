package elements

// The Data element contains encrypted data that represents the shared data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/data
import "encoding/xml"

type Data struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *Data) SetForMarshal() {
	D.XMLName.Local = "t:Data"
}

func (D *Data) GetSchema() *Schema {
	return &SchemaTypes
}
