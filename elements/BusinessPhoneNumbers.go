package elements

// The BusinessPhoneNumbers element specifies an array of business phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businessphonenumbers
import "encoding/xml"

type BusinessPhoneNumbers struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies a phone number and type information and is associated with a set of attributions.
	Value *ValuePersonaPhoneNumberType `xml:"Value"`
}

func (B *BusinessPhoneNumbers) SetForMarshal() {
	B.XMLName.Local = "t:BusinessPhoneNumbers"
}

func (B *BusinessPhoneNumbers) GetSchema() *Schema {
	return &SchemaTypes
}
