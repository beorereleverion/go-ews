package elements

// The CallbackPhones element specifies an array of call-back phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/callbackphones
import "encoding/xml"

type CallbackPhones struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies a phone number and type information and is associated with a set of attributions.
	Value *ValuePersonaPhoneNumberType `xml:"Value"`
}

func (C *CallbackPhones) SetForMarshal() {
	C.XMLName.Local = "t:CallbackPhones"
}

func (C *CallbackPhones) GetSchema() *Schema {
	return &SchemaTypes
}
