package elements

// The RadioPhones element specifies an array of radio phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/radiophones
import "encoding/xml"

type RadioPhones struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (R *RadioPhones) SetForMarshal() {
	R.XMLName.Local = "t:RadioPhones"
}

func (R *RadioPhones) GetSchema() *Schema {
	return &SchemaTypes
}
