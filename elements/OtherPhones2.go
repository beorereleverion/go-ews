package elements

// The OtherPhones2 element specifies an array of phone values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/otherphones2
import "encoding/xml"

type OtherPhones2 struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (O *OtherPhones2) SetForMarshal() {
	O.XMLName.Local = "t:OtherPhones2"
}

func (O *OtherPhones2) GetSchema() *Schema {
	return &SchemaTypes
}
