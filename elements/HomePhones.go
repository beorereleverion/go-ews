package elements

// The HomePhones element specifies an array of home phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homephones
import "encoding/xml"

type HomePhones struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (H *HomePhones) SetForMarshal() {
	H.XMLName.Local = "t:HomePhones"
}

func (H *HomePhones) GetSchema() *Schema {
	return &SchemaTypes
}
