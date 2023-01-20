package elements

// The OriginalPhoneString element specifies the original phone number for a contact or persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalphonestring
import "encoding/xml"

type OriginalPhoneString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (O *OriginalPhoneString) SetForMarshal() {
	O.XMLName.Local = "t:OriginalPhoneString"
}

func (O *OriginalPhoneString) GetSchema() *Schema {
	return &SchemaTypes
}
