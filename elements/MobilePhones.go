package elements

// The MobilePhones element specifies an array of mobile phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mobilephones
import "encoding/xml"

type MobilePhones struct {
	XMLName xml.Name

	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

func (M *MobilePhones) SetForMarshal() {
	M.XMLName.Local = "t:MobilePhones"
}

func (M *MobilePhones) GetSchema() *Schema {
	return &SchemaTypes
}
