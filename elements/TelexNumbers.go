package elements

// The TelexNumbers element specifies an array of Telex numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/telexnumbers
import "encoding/xml"

type TelexNumbers struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (T *TelexNumbers) SetForMarshal() {
	T.XMLName.Local = "t:TelexNumbers"
}

func (T *TelexNumbers) GetSchema() *Schema {
	return &SchemaTypes
}
