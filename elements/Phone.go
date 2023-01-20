package elements

// The Phone element specifies a single phone number that results from a contact entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phone
import "encoding/xml"

type Phone struct {
	XMLName xml.Name

	// The OriginalPhoneString element specifies the original phone number for a contact or persona.
	OriginalPhoneString *OriginalPhoneString `xml:"OriginalPhoneString"`
	// The PhoneString element specifies the phone number for an extracted contact.
	PhoneString *PhoneString `xml:"PhoneString"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"Type"`
}

func (P *Phone) SetForMarshal() {
	P.XMLName.Local = "t:Phone"
}

func (P *Phone) GetSchema() *Schema {
	return &SchemaTypes
}
