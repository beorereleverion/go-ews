package elements

// The ImTelephoneNumber element represents the telephone number for a contact that is added to an instant messaging (IM) group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imtelephonenumber
import "encoding/xml"

type ImTelephoneNumber struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ImTelephoneNumber) SetForMarshal() {
	I.XMLName.Local = "m:ImTelephoneNumber"
}

func (I *ImTelephoneNumber) GetSchema() *Schema {
	return &SchemaMessages
}
