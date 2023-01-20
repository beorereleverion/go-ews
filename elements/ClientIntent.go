package elements

// The ClientIntent element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/clientintent
import "encoding/xml"

type ClientIntent struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (C *ClientIntent) SetForMarshal() {
	C.XMLName.Local = "m:ClientIntent"
}

func (C *ClientIntent) GetSchema() *Schema {
	return &SchemaMessages
}
