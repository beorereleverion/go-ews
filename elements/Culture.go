package elements

// The Culture element represents the culture for a given item in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/culture
import "encoding/xml"

type Culture struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *Culture) SetForMarshal() {
	C.XMLName.Local = "t:Culture"
}

func (C *Culture) GetSchema() *Schema {
	return &SchemaTypes
}
