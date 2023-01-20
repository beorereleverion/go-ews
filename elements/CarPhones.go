package elements

// The CarPhone element specifies an array of car phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/carphones
import "encoding/xml"

type CarPhones struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies a phone number and type information and is associated with a set of attributions.
	Value *ValuePersonaPhoneNumberType `xml:"Value"`
}

func (C *CarPhones) SetForMarshal() {
	C.XMLName.Local = "t:CarPhones"
}

func (C *CarPhones) GetSchema() *Schema {
	return &SchemaTypes
}
