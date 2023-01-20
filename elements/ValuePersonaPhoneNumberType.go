package elements

// The Value element specifies a phone number and type information and is associated with a set of attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-personaphonenumbertype
import "encoding/xml"

type ValuePersonaPhoneNumberType struct {
	XMLName xml.Name

	// The Number element specifies a phone number.
	Number *Number `xml:"Number"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"Type"`
}

func (V *ValuePersonaPhoneNumberType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValuePersonaPhoneNumberType) GetSchema() *Schema {
	return &SchemaTypes
}
