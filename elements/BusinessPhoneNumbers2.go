package elements

// The BusinessPhoneNumbers2 element specifies an array of BusinessPhoneNumber2 elements and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businessphonenumbers2
import "encoding/xml"

type BusinessPhoneNumbers2 struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies a phone number and type information and is associated with a set of attributions.
	Value *ValuePersonaPhoneNumberType `xml:"Value"`
}

func (B *BusinessPhoneNumbers2) SetForMarshal() {
	B.XMLName.Local = "t:BusinessPhoneNumbers2"
}

func (B *BusinessPhoneNumbers2) GetSchema() *Schema {
	return &SchemaTypes
}
