package elements

// The PhoneNumber element specifies the default phone number of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumber
import "encoding/xml"

type PhoneNumber struct {
	XMLName xml.Name

	// The Number element specifies a phone number.
	Number *Number `xml:"Number"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"Type"`
}

func (P *PhoneNumber) SetForMarshal() {
	P.XMLName.Local = "t:PhoneNumber"
}

func (P *PhoneNumber) GetSchema() *Schema {
	return &SchemaTypes
}
