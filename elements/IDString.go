package elements

// The ID element specifies the identifier of an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/id-string
import "encoding/xml"

type IDString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *IDString) SetForMarshal() {
	I.XMLName.Local = "m:ID"
}

func (I *IDString) GetSchema() *Schema {
	return &SchemaMessages
}
