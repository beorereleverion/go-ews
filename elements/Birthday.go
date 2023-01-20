package elements

// The Birthday element represents the birth date of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/birthday
import "time"

import "encoding/xml"

type Birthday struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (B *Birthday) SetForMarshal() {
	B.XMLName.Local = "t:Birthday"
}

func (B *Birthday) GetSchema() *Schema {
	return &SchemaTypes
}
