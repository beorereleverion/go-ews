package elements

// The HomePhones2 element specifies an array of HomePhone2 values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homephones2
import "encoding/xml"

type HomePhones2 struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (H *HomePhones2) SetForMarshal() {
	H.XMLName.Local = "t:HomePhones2"
}

func (H *HomePhones2) GetSchema() *Schema {
	return &SchemaTypes
}
