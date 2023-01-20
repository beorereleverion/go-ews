package elements

// The TTYTDDPhoneNumbers element specifies an array of TTY or TDD text telephone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ttytddphonenumbers
import "encoding/xml"

type TTYTDDPhoneNumbers struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (T *TTYTDDPhoneNumbers) SetForMarshal() {
	T.XMLName.Local = "t:TTYTDDPhoneNumbers"
}

func (T *TTYTDDPhoneNumbers) GetSchema() *Schema {
	return &SchemaTypes
}
