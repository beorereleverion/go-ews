package elements

// The AssistantPhoneNumbers element specifies an array of assistant phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assistantphonenumbers
import "encoding/xml"

type AssistantPhoneNumbers struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (A *AssistantPhoneNumbers) SetForMarshal() {
	A.XMLName.Local = "t:AssistantPhoneNumbers"
}

func (A *AssistantPhoneNumbers) GetSchema() *Schema {
	return &SchemaTypes
}
