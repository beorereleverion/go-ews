package elements

// The PhoneString element specifies the phone number for an extracted contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonestring
import "encoding/xml"

type PhoneString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PhoneString) SetForMarshal() {
	P.XMLName.Local = "t:PhoneString"
}

func (P *PhoneString) GetSchema() *Schema {
	return &SchemaTypes
}
