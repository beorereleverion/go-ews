package elements

// The Phone element specifies a single phone number that results from a phone number entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phone-phoneentitytype
import "encoding/xml"

type PhonePhoneEntityType struct {
	XMLName xml.Name

	// The OriginalPhoneString element specifies the original phone number for a contact or persona.
	OriginalPhoneString *OriginalPhoneString `xml:"OriginalPhoneString"`
	// The PhoneString element specifies the phone number for an extracted contact.
	PhoneString *PhoneString `xml:"PhoneString"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"Position"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"Type"`
}

func (P *PhonePhoneEntityType) SetForMarshal() {
	P.XMLName.Local = "t:Phone"
}

func (P *PhonePhoneEntityType) GetSchema() *Schema {
	return &SchemaTypes
}
