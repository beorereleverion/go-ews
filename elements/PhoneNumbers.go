package elements

// The PhoneNumbers element represents a collection of telephone numbers for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers
import "encoding/xml"

type PhoneNumbers struct {
	XMLName xml.Name

	// The Entry element represents a telephone number for a contact.
	Entry *EntryPhoneNumber `xml:"Entry"`
}

func (P *PhoneNumbers) SetForMarshal() {
	P.XMLName.Local = "t:PhoneNumbers"
}

func (P *PhoneNumbers) GetSchema() *Schema {
	return &SchemaTypes
}
