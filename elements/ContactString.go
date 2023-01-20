package elements

// The ContactString element specifies the display name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactstring
import "encoding/xml"

type ContactString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContactString) SetForMarshal() {
	C.XMLName.Local = "t:ContactString"
}

func (C *ContactString) GetSchema() *Schema {
	return &SchemaTypes
}
